import { dev } from '$app/environment';
import type { Folder } from '../models';
import type PaginationData from '../models/common/pagination-data';

const apiPrefix = (dev ? 'http://localhost:8080' : '') + '/api';
const api = {
	folders: `${apiPrefix}/folders`
};

export const getFolders = async (
	fetchFn: typeof fetch,
	page: number,
	size: number | null = null
): Promise<{ records: Folder[]; pagination: PaginationData } | null> => {
	try {
		return await (await fetchFn(`${api.folders}?current=${page}&size=${size || 10}`)).json();
	} catch (err) {
		console.log(err);
		return null;
	}
};
