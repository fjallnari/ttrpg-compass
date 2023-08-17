export default interface TTRPGSystem {
    Id: string;
    Title: string;
    Edition: string;
    Description: string;
    Publisher: string;

    Complexity: number;
    Progression: number;
    Narrative: number;
    Combat: number;
    Exploration: number;
    Balance: number;
    Versatility: number;
    Customization: number;

    Genre: string;
    Family: string;
    Gm: string;
    Similar: string[];
}