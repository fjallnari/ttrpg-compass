import os

def delete_systems(dir_name):
    dir = os.listdir(dir_name)

    for file in dir:
        if file.endswith(".toml"):
            os.remove(os.path.join(dir_name, file))

if __name__ == "__main__":
    delete_systems("../stress_test")
