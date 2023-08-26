import { writable, type Writable } from "svelte/store";
import type TTRPGSystem from "./interfaces/TTRPGSystem";

export const serverURL: Writable<string> = writable("");
export const cursor: Writable<number> = writable(0);
export const foundSystems: Writable<TTRPGSystem[]> = writable([]);
export const selectedSystem: Writable<TTRPGSystem | undefined> = writable(undefined);