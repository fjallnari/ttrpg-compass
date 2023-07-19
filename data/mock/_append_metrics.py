import os
from random import randint

def append_metrics(folder):
    files = os.listdir(folder)
    for file_path in files:
        if file_path.endswith('.toml'):
            with open(file_path, "a") as file:
                file.write("\n")
                for metric in ["complexity", "progression", "narrative", "combat", "exploration",  "balance", "versatility", "customization"]:
                    file.write(f"\n{metric} = {randint(1, 5)}")
                
append_metrics(".")