from random import choice
from nanoid import generate

adjectives = [
    "Astral",
    "Celestial",
    "Cosmic",
    "Eldritch",
    "Voyager's",
    "Arcane",
    "Enigma",
    "Darkmoon",
    "Starward",
    "Mystic",
    "Epic",
    "Terra",
    "Eclipse",
    "Mythica",
    "Nexus",
    "Enchanted",
    "Elemental",
    "Stellar",
    "Forgotten",
    "Stardrift",  
    "Lush",
    "Eco",
    "Sunlit",
    "Rustland",
    "Dystopian",
    "Nuclear",
    "Ashen",
    "Endless"
]

preposition_noun = [
    "of the Celestials",
    "of the Void",
    "of the Aether",
    "of the Stars",
    "of Nevermore",
    "of Yesteryear",
    "of Starlight",
    "of the Eons",
    "of Stardust",
    "of the Cosmos",
    "of Avendel",
    "of Gaia",
    "of the Abyss",
    "of Evermore",
    "of the Eternals",
    "of Eldoria",
    "of the Empyrean",
    "of Mythos",
    "of Lumina",
    "of Eldarune",
    "of Elara",
    "of Astraeus",
]

nouns = [
    "Ascendancy",
    "Realms",
    "Song",
    "Chronicles",
    "Prophecy",
    "Voyages",
    "Aetheria",
    "Sagas",
    "Empires",
    "Legends",
    "Horizons",
    "Tales",
    "Odyssey",
    "Mystica",
    "Terra",
    "Infinity",
    "Journeys",
    "Eternity",
    "Quest",
    "Haven",
    "Elysium",
    "Aurora",
    "Light",
    "Grove",
    "Darkness",
    "Shadow",
    "Dawn",
    "Dusk",
    "Twilight",
    "Rebirth",
    "Utopia",
    "Requiem",
    "Runners",
    "Spirits",
    "Revolution",
    "Revelation",
    "Incursion",
    "Genesis",
    "Empire",
    "Protocol",
    "Insurgency",
    "Matrix",
    "Outrunners",
    "Outriders",
    "Expanse",
    "Exodus",
    "Dominion"
]

compound_start = [
    "Lore",
    "Star",
    "Warp",
    "Ever"
    "World",
    "Myth",
    "Realm",
    "Never",
    "Silk",
    "Chain",
    "Soul",
    "Cyber",
    "Steam",
    "Solar",
    "Clock",
    "Synth",
    "Bio",
    "Neuro",
    "Vita",
    "Pixel",
    "Arc",
    "Ruin",
    "Rune",
    "Rift",
    "Ashen"
]

compound_end = [
    "borne",
    "wind",
    "ward"
    "weave"
    "weaver",
    "walker",
    "forge",
    "strider",
    "maker",
    "seeker",
    "binder",
    "bound",
    "breaker",
    "walker",
    "fall",
    "pulse",
    "wave",
    "drift",
    "tech",
    "dyne",
]

genres = [
    "Dark Fantasy",
    "Mythic Fantasy",
    "Sword & Sorcery",
    "Heroic Fantasy"
    "Sci-Fi",
    "Space Opera",
    "Cyberpunk",
    "Steampunk",
    "Post-Apocalyptic",
    "Horror",
    "Supernatural",
    "Mystery",
    "Historical",
    "Western",
    "Superhero",
    "Cosmic Horror",
    "Grimdark",
    "Viking",
    "Mechs",
    "Pirates",
    "Solarpunk",
    "Clockpunk",
    "Dieselpunk",
    "Biopunk",
    "Obscure",
    "Weird West",
    "Who knows?"
]

system_families = [
    "D&D",
    "OSR",
    "NSR",
    "PbtA",
    "Cypher System",
    "Fate",
    "Savage Worlds",
    "GURPS",
    "BRP",
    "Fudge",
    "D6",
    "D20",
    "Card-based",
    "Shadowrun",
    "Warhammer",
    "Indie",
    "GUMSHOE",
]


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
        self.type = choice(system_families)
        self.genre = choice(genres)
        self.edition = self.gen_edition()
        self.gm = choice(20 * ["Game Master"] + ["GM-less/Solo", "Shared GM"])
        self.description = "Venture into a world where elemental forces shape the land, and heroes rise to challenge ancient dragons and mythical beasts."
        self.system = {'title': self.title, 'type': self.type, 'genre': self.genre, 'edition': self.edition, 'gm': self.gm, 'description': self.description}
        
    def gen_edition(self):
        return f'{choice(["1st", "2nd", "3rd", "4th", "5th", "6th", "7th", "8th"])} edition'
    
    def get_system(self):
        return self.system


def generate_n_systems(n):
    for _ in range(n):
        id = generate('0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz', 8)
        system = SystemGenerator().get_system()
        with open(f"{id}.toml", "a+") as file:
            file.write(f"title = \"{system['title']}\"\ndescription = \"{system['description']}\"\ntype = \"{system['type']}\"\ngenre = \"{system['genre']}\"\nedition = \"{system['edition']}\"\ngm = \"{system['gm']}\"\n")

generate_n_systems(200)