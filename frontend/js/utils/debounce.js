/**
 * Creates a debounced function that delays invoking `func`
 * until after `wait` milliseconds have elapsed since the
 * last time the debounced function was invoked.
 */
export function debounce(func, wait) {
    let timeoutId = null;

    const debounced = function (...args) {
        clearTimeout(timeoutId);

        timeoutId = setTimeout(() => {
            func.apply(this, args);
        }, wait);
    };

    debounced.cancel = function () {
        clearTimeout(timeoutId);
        timeoutId = null;
    };

    return debounced;
}
