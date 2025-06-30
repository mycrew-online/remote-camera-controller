<script lang="ts">
	import '../app.css';
	
	let { children } = $props();

	// WebSocket connection to backend
	if (typeof window !== 'undefined') {
		console.log('Initializing WebSocket connection...');
		const socket = new WebSocket('ws://localhost:8080/ws');

		socket.onopen = () => {
			console.log('WebSocket connected');
			// Optionally send a test message:
			// socket.send('Hello from SvelteKit!');
		};

		socket.onmessage = (event) => {
			console.log('Message from backend:', event.data);
		};

		socket.onerror = (error) => {
			console.error('WebSocket error:', error);
		};

		socket.onclose = () => {
			console.log('WebSocket closed');
		};
	}
</script>

{@render children()}
