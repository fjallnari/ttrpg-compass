import tomli
import os
import numpy as np
from metrics_names import METRICS_NAMES


def get_metrics_as_array(toml_dict):
    metrics = []
    
    for metric in METRICS_NAMES:
        metrics.append(toml_dict[metric])
        
    return metrics

def load_toml_systems(path):
    dir = os.listdir(path)
    systems = {}
    
    for file_name in dir:
        if file_name.endswith(".toml"):
            with open(f"{path}/{file_name}", "rb") as f:
                toml_dict = tomli.load(f)
                # print(file_name.split('.')[0], get_metrics_as_array(toml_dict), toml_dict['genre'], toml_dict['family'])
                systems[file_name.split('.')[0]] = {
                    'genre': toml_dict['genre'],
                    'family': toml_dict['family'],
                    'metrics': get_metrics_as_array(toml_dict),
                }
    return systems

def calc_similarity(systemA, systemB, genreWeight=3, familyWeight=3):
    sameGenre = genreWeight if systemA['genre'] == systemB['genre'] else 0
    sameFamily = familyWeight if systemA['family'] == systemB['family'] else 0
    
    point1 = np.array(systemA['metrics'])
    point2 = np.array(systemB['metrics'])
    
    # print(np.linalg.norm(point1 - point2) - sameGenre - sameFamily) 
    return np.linalg.norm(point1 - point2) - sameGenre - sameFamily
    
def calc_all_similarities(systems):
    similarities = {}
    
    for systemA in systems:
        similarities[systemA] = {}
        
        for systemB in systems:
            if systemA != systemB:
                # print(f"{systemA} ~ {systemB}: ", end="")
                similarities[systemA][systemB] = calc_similarity(systems[systemA], systems[systemB])
            
    return similarities

def append_similar_to_toml(most_similar, path):
    for system in most_similar:
        with open(f"{path}/{system}.toml", "a+") as file:
            file.write(f"\n\nsimilar = {most_similar[system]}")

def get_3_most_similar(similarities):
    most_similar = {}
    
    for system in similarities:
        most_similar[system] = sorted(similarities[system], key=similarities[system].get)[:3]
        
    return most_similar

def calc_and_write_similarities(path):
    systems = load_toml_systems(path)
    similarities = calc_all_similarities(systems)
    most_similar = get_3_most_similar(similarities)
    
    append_similar_to_toml(most_similar, path)

if __name__ == "__main__":
    calc_and_write_similarities("../stress_test")