export default class PaginationData {
	current: number;
	size: number;
	totalItems: number;
	totalPages: number;

	constructor(result: any) {
		this.current = result.current;
		this.size = result.size;
		this.totalItems = result.totalItems;
		this.totalPages = Math.ceil(this.totalItems / this.size);
	}

	get isVisible() {
		return this.totalPages > 1;
	}

	get isNextEnabled() {
		return this.current < this.totalPages;
	}

	get isPreviousEnabled() {
		return this.current > 1;
	}
}
