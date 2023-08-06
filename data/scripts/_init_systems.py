from _gen_systems import generate_n_systems
from _append_metrics import append_metrics
from _calc_similarity import calc_and_write_similarities

def init_systems(n, path):
    generate_n_systems(n, path)
    append_metrics(path)
    calc_and_write_similarities(path)
    print('OK')
    
if __name__ == "__main__":
    init_systems(5, "../stress_test")