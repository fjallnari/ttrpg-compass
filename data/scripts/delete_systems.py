import os
import sys

def delete_systems(dir_name, flushdb):
    dir = os.listdir(dir_name)

    for file in dir:
        if file.endswith(".toml"):
            os.remove(os.path.join(dir_name, file))
    
    if flushdb:
        os.system("docker exec -it redis-stack redis-cli FLUSHALL")

if __name__ == "__main__":
    flushdb = sys.argv[1] == '-f' if len(sys.argv) == 2 else False
    delete_systems("../stress_test", flushdb)
