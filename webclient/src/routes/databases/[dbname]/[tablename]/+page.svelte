<script>
	/** @type {import('./$types').PageData} */
	export let data;
</script>

<!-- component -->
<!-- This is an example component -->
<div class="max-w-2xl mx-auto mt-10">
	<div class="relative shadow-md sm:rounded-lg">
		<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
			<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 ">
				<tr >
					{#each data.table.headers as header}
						<th scope="col" class="px-6 py-3">
							<div>{header.name}</div>
							<br />
							<div class="text-smol">
								{header.type}
							</div>
						</th>
					{/each}
					<th scope="col" class="px-6 py-3">
						<span class="sr-only">Edit</span>
					</th>
				</tr>
			</thead>
			<tbody>
				{#each data.table.values as row}
					<tr
						class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
					>
						{#each row as value}
							<td class="px-6 py-4"> {value} </td>
						{/each}
						<td class="px-6 py-4 text-right">
							<a href="#" class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
								>Edit</a
							>
						</td>
					</tr>
				{/each}
                
			</tbody>
		</table>
	</div>

	<script src="https://unpkg.com/flowbite@1.3.4/dist/flowbite.js"></script>
</div>

<style>
	.text-smol {
		font-size: 0.6rem;
		line-height: 1rem;
	}
</style>
