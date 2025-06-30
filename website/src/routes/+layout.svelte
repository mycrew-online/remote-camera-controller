<script lang="ts">
	import '../app.css';
	
	let { children } = $props();

	// WebSocket connection to backend
	if (typeof window !== 'undefined') {
		console.log('Initializing WebSocket connection...');
		const socket = new WebSocket('ws://localhost:8080/ws');

		socket.onopen = () => {
			console.log('WebSocket connected');
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

<div class="min-h-screen flex flex-col bg-gray-50 dark:bg-gray-900">
	<header class="w-full px-4 py-3 bg-blue-600 text-white shadow-md flex items-center justify-between">
		<h1 class="text-lg font-bold tracking-wide">Remote Camera Controller</h1>
		<span class="text-xs opacity-80">v1.0</span>
	</header>
	<main class="flex-1 flex flex-col items-center justify-start p-4 w-full max-w-lg mx-auto">
		{@render children()}
	</main>
	<footer class="w-full py-2 text-center text-xs text-gray-400 bg-transparent mt-auto">
		&copy; {new Date().getFullYear()} MyCrew.online
	</footer>
</div>
