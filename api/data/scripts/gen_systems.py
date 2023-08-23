from random import choice
from nanoid import generate
from gen_data import adjectives, nouns, compound_start, compound_end, system_families, genres, publishers, preposition_noun

class NameGenerator:    
    def adj_noun(self):
        return f"{choice(adjectives)} {choice(nouns)}"
    
    def compound(self):
        return f"{choice(compound_start)}{choice(compound_end)}"
    
    def preposition_noun(self):
        return f"{choice(nouns)} {choice(preposition_noun)}"
    
    def get_name(self):
        return choice([self.adj_noun, self.compound, self.preposition_noun])()

class SystemGenerator:
    def __init__(self):
        self.title = NameGenerator().get_name()
        self.family = choice(system_families)
        genre = choice(genres)
        self.genre = genre["name"].lower()
        self.edition = self.gen_edition()
        self.publisher = choice(publishers)
        self.gm = choice(20 * ["Game Master"] + ["GM-less/Solo", "Shared GM"])
        self.description = genre["description"]
        self.system = {
            'title': self.title, 
            'family': self.family, 
            'genre': self.genre, 
            'edition': self.edition, 
            'gm': self.gm, 
            'description': self.description, 
            'publisher': self.publisher
        }
        
    def gen_edition(self):
        return f'{choice(["1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th"])} edition'
    
    def get_system(self):
        return self.system


def generate_n_systems(n, path):
    for _ in range(n):
        id = generate('0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz', 8)
        system = SystemGenerator().get_system()
        with open(f"{path}/{id}.toml", "a+") as file:
            file.write(f"title = \"{system['title']}\"\ndescription = \"{system['description']}\"\nfamily = \"{system['family']}\"\ngenre = \"{system['genre']}\"\nedition = \"{system['edition']}\"\ngm = \"{system['gm']}\"\npublisher = \"{system['publisher']}\"\n")

if __name__ == "__main__":
    generate_n_systems(1, ".")