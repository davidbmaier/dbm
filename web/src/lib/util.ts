import { PUBLIC_API_URL } from '$env/static/public';
import type { Work } from '../types';

export const getFileLink = (work: Work) => {
	return `${PUBLIC_API_URL}/files/${work.workID}`;
};

export const updateStorage = (entries: { id: string; value: number | string }[]) => {
	for (const entry of entries) {
		localStorage.setItem(entry.id, `${entry.value}`);
	}
};

export interface storageEntry {
	[id: string]: string | null;
}
export const getStorage = (entries: string[]): storageEntry => {
	const resolvedEntries: storageEntry = {};
	for (const entry of entries) {
		resolvedEntries[entry] = localStorage.getItem(entry);
	}
	return resolvedEntries;
};

export const debounce = <Params extends unknown[]>(
	func: (...args: Params) => unknown,
	timeout = 400
): ((...args: Params) => void) => {
	let timer: number;
	return (...args: Params) => {
		clearTimeout(timer);
		timer = setTimeout(() => {
			func(...args);
		}, timeout);
	};
};

export const getTitle = (titleContent: string) => {
	return titleContent ? `${titleContent} - DBM` : ``;
};
