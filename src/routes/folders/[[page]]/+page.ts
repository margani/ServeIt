import type { PageLoad } from './$types';
import { FoldersService } from '../../../services';
import PaginationData from '../../../models/common/pagination-data';
import { loading } from '$lib'

export const load: PageLoad = async ({ fetch, params }) => {
	loading.starting();
	const page: number = parseInt(params.page || '1');
	const data = await FoldersService.getFolders(fetch, page);
	return {
		records: data?.records || [],
		pagination: new PaginationData(data?.pagination || { current: 1, size: 10, totalItems: 0 })
	};
};
