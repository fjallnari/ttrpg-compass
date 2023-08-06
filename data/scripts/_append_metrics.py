import os
from random import randint

def append_metrics(path):
    files = os.listdir(path)
    for file_name in files:
        if file_name.endswith('.toml'):
            with open(f"{path}/{file_name}", "a") as file:
                for metric in ["complexity", "progression", "narrative", "combat", "exploration",  "balance", "versatility", "customization"]:
                    file.write(f"\n{metric} = {randint(1, 5)}")

if __name__ == "__main__":
    append_metrics(".")
