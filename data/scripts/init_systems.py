import sys

from gen_systems import generate_n_systems
from append_metrics import append_metrics
from calc_similarity import calc_and_write_similarities

def init_systems(n, path):
    generate_n_systems(n, path)
    append_metrics(path)
    calc_and_write_similarities(path)
    print('OK')
    
if __name__ == "__main__":
    no_of_systems = 5 if len(sys.argv) == 1 else int(sys.argv[1])
    init_systems(no_of_systems, "../stress_test")