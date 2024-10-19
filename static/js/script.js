document.addEventListener('alpine:init', () => {
    Alpine.data('counter', () => {
        return {
            count: 1,
            
            increment() {
                this.count++;
            },
        }
    })
})