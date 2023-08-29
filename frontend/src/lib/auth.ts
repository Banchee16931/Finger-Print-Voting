import { writable } from 'svelte/store';
import type { UserData } from './types';

export const user = writable<UserData>();