<script>
	/** @type {import('./$types').PageData} */
	export let data;

	import { PUBLIC_API_BASE_URL } from '$env/static/public';
	import { invalidateAll } from '$app/navigation';
	import axios from 'axios';
	import { each } from 'svelte/internal';

	let select1, select2 
	
</script>

<label for="s1" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select first table</label>
<select bind:this={select1} id="s1" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
	<option selected>Choose a table 1</option>
	{#each data.info.tables as table, i}
		<option value="{i}">{table.name}</option>
	{/each}
</select>

<label for="s2" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select second table</label>
<select bind:this={select2} id="s2" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
	<option selected>Choose a table 2</option>
	{#each data.info.tables as table, i}
		<option value="{i}">{table.name}</option>
	{/each}
</select>

<label for="s21" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Select first column</label>
<select bind:this={select1} id="s21" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
	<option selected>Choose column 1</option>
	{#each data.info.tables[0].headers as header, i}
		<option value="{i}">{header.name} :: {header.type}</option>
	{/each}
</select>