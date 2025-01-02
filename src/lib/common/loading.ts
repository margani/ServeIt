import { writable } from 'svelte/store';

const newLoading = () => {
    const { subscribe, update, set } = writable({
        status: 'IDLE', // IDLE, LOADING, NAVIGATING
    });

    function setNavigate(isNavigating: boolean) {
        update(() => {
            return {
                status: isNavigating ? 'NAVIGATING' : 'IDLE',
                message: '',
            };
        });
    }

    function finished() {
        update(() => {
            return {
                status: 'IDLE',
            };
        });
    }

    function starting() {
        update(() => {
            return {
                status: 'LOADING',
            };
        });
    }

    return { subscribe, update, set, setNavigate, starting, finished, };
};

export const loading = newLoading();
