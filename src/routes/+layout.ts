export const prerender = false;
import folders from '../store/folders';

export async function load({ fetch }) {
	await folders.fetch(fetch);

	return {
		props: {}
	};
}
