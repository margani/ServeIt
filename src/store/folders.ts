import { writable } from 'svelte/store';
import { FoldersService } from '../services';
import type { Folder } from '../models';

function createFoldersStore() {
	const { subscribe, set, update } = writable([] as Folder[]);

	return {
		subscribe,
		set: (folders: Folder[]) => set(folders),
		fetch: async (fetchFn: typeof fetch) => {
			const data = await FoldersService.getFolders(fetchFn, 1, 100);
			set(data?.records || []);
		}
	};
}

const folders = createFoldersStore();

export default folders;
