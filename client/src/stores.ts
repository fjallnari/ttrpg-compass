import { writable, type Writable } from "svelte/store";

export const cursor: Writable<number> = writable(0);