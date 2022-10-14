<script>
	export let data;
</script>

<div class="sticky top-0 flex flex-col justify-center overflow-hidden bg-gray-50">
	<div class="flex items-center justify-between bg-blue-500 py-7">
		<div class="absolute left-1/2 -translate-x-1/2 text-lg font-bold text-gray-100">
			😎 Database Management System 😎
		</div>
	</div>
</div>

<div class="container mx-auto flex items-center justify-center flex-col mt-10">
	<div class="min-w-[50%] w-full sm:w-auto">
		<div class="flex flex-col justify-center">
			<div class="overflow-x-auto shadow-md sm:rounded-lg">
				<div class="inline-block min-w-full align-middle">
					<div class="overflow-hidden">
						<table class="min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-700">
							<tbody
								class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700"
							>
								{#each data.databases as database}
									<tr class="relative hover:bg-gray-100 dark:hover:bg-gray-700">
										<td
											class="py-7 px-6 text-sm  font-medium text-center text-gray-900 whitespace-nowrap dark:text-white"
										>
											<a
												class="absolute left-0 top-0 flex items-center justify-center w-full h-full"
												href="/databases/{database}">{database}</a
											>
										</td>
									</tr>
								{/each}
							</tbody>
						</table>
					</div>
				</div>
			</div>
			<button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
				Create new database
			</button>
		</div>
	</div>
</div>

<style>
</style>
