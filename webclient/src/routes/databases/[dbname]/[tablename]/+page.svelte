<script>
	/** @type {import('./$types').PageData} */
	export let data;

	import { PUBLIC_API_BASE_URL } from '$env/static/public';
	import { invalidateAll } from '$app/navigation';
	import axios from 'axios';
	import { each } from 'svelte/internal';

	let addData = [];
	let editData = [];
	let boba = 'ddd';
	let url = `${PUBLIC_API_BASE_URL}/databases/${data.dbname}/${data.tablename}`;
	let modalEditView;
	let selectedID;

	function onAdd() {
		let addUrl = `${url}/new_row`;
		axios
			.post(addUrl, addData)
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});

		addData = [];

		invalidateAll();
	}

	function onDelete(id) {
		let delUrl = `${url}/${id}`;

		axios
			.delete(delUrl)
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});

		invalidateAll();
	}

	function onEdit() {
		let editUrl = `${url}/${selectedID}`;

		axios
			.put(editUrl, editData.map(String))
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});

		editData = [];
		fadeOut(modalEditView);
		invalidateAll();
	}

	function modalHandler(val, i) {
		selectedID = i;
		if (val) {
			fadeIn(modalEditView);
		} else {
			fadeOut(modalEditView);
		}
	}

	function fadeOut(el) {
		el.style.opacity = 1;
		(function fade() {
			if ((el.style.opacity -= 0.1) < 0) {
				el.style.display = 'none';
			} else {
				requestAnimationFrame(fade);
			}
		})();
	}

	function fadeIn(el, display) {
		data.table.values[selectedID].forEach(function (value, i) {
			editData[i] = value
		});

		el.style.opacity = 0;
		el.style.display = display || 'flex';
		(function fade() {
			let val = parseFloat(el.style.opacity);
			if (!((val += 0.2) > 1)) {
				el.style.opacity = val;
				requestAnimationFrame(fade);
			}
		})();
	}
</script>

<!--  -->
<!-- edit -->
<dh-component>
	<div
		class="py-20 transition duration-150 ease-in-out z-10 absolute top-0 right-0 bottom-0 left-0"
		style="display:none"
		bind:this={modalEditView}
	>
		<div role="alert" class="container mx-auto w-11/12 md:w-2/3 max-w-lg">
			<div class="relative py-8 px-5 md:px-10 bg-white shadow-md rounded border border-gray-400">
				<h1 class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4">Edit row</h1>

				{#each data.table.headers as header, j}
					<label for="name" class="text-gray-800 text-sm font-bold leading-tight tracking-normal"
						>{header.name}</label
					>
					<input
						id="name"
						type="text"
						class="mb-5 mt-2 text-gray-600 focus:outline-none focus:border focus:border-indigo-700 font-normal w-full h-10 flex items-center pl-3 text-sm border-gray-300 rounded border"
						bind:value={editData[j]}
					/>
				{/each}

				<div class="flex items-center justify-start w-full">
					<button
						class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-700 transition duration-150 ease-in-out hover:bg-indigo-600 bg-indigo-700 rounded text-white px-8 py-2 text-sm"
						on:click={onEdit}
						>Submit</button
					>
					<button
						class="focus:outline-none focus:ring-2 focus:ring-offset-2  focus:ring-gray-400 ml-3 bg-gray-100 transition duration-150 text-gray-600 ease-in-out hover:border-gray-400 hover:bg-gray-300 border rounded px-8 py-2 text-sm"
						on:click|preventDefault={() => modalHandler(false, -1)}>Cancel</button
					>
				</div>
				<button
					class="cursor-pointer absolute top-0 right-0 mt-4 mr-5 text-gray-400 hover:text-gray-600 transition duration-150 ease-in-out rounded focus:ring-2 focus:outline-none focus:ring-gray-600"
					on:click|preventDefault={() => modalHandler(false, -1)}
					aria-label="close modal"
					role="button"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="icon icon-tabler icon-tabler-x"
						width="20"
						height="20"
						viewBox="0 0 24 24"
						stroke-width="2.5"
						stroke="currentColor"
						fill="none"
						stroke-linecap="round"
						stroke-linejoin="round"
					>
						<path stroke="none" d="M0 0h24v24H0z" />
						<line x1="18" y1="6" x2="6" y2="18" />
						<line x1="6" y1="6" x2="18" y2="18" />
					</svg>
				</button>
			</div>
		</div>
	</div>
</dh-component>

<!-- header -->
<div class="sticky top-0 flex flex-col justify-center overflow-hidden bg-gray-50">
	<div class="flex items-center justify-between bg-blue-500 p-3">
		<a
			href="/databases/{data.dbname}"
			class="flex items-center space-x-2 rounded bg-gray-100 py-1 px-2 text-slate-500 shadow-md hover:bg-white"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				class="h-4 w-4"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M19.5 12h-15m0 0l6.75 6.75M4.5 12l6.75-6.75"
				/>
			</svg>
			<span>Back</span>
		</a>
		<div class="absolute left-1/2 -translate-x-1/2 text-lg font-bold text-gray-100">
			{data.dbname}/{data.tablename}
		</div>
	</div>
</div>

<!-- component -->
<!-- This is an example component -->
<div class="max-w-2xl mx-auto mt-10">
	<div class="relative shadow-md sm:rounded-lg">
		<table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
			<thead
				class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 "
			>
				<tr>
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
					<th />
				</tr>

				<tr class="bg-gray-700 border-t-2 border-gray-600">
					{#each data.table.headers as header, i}
						<th scope="col" class="px-6 py-3">
							<input
								type="text"
								id="last_name"
								bind:value={addData[i]}
								class="text-xs rounded-lg   block w-full p-2.5 bg-gray-800 border-gray-600 placeholder-gray-500 text-white outline-none"
								placeholder={header.name}
								required
							/>
						</th>
					{/each}
					<td class="px-6 py-4 text-right">
						<a
							href="#"
							class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
							on:click={onAdd}>Add</a
						>
					</td>
					<td />
				</tr>
			</thead>
			<tbody>
				{#each data.table.values as row, i}
					<tr
						class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
					>
						{#each row as value}
							<td class="px-6 py-4"> {value} </td>
						{/each}
						<td class="px-6 py-4 text-right">
							<a
								href="#"
								class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
								on:click|preventDefault={() => modalHandler(true, i)}>Edit</a
							>
						</td>
						<td class="px-6 py-4 text-right">
							<a
								href="#"
								class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
								on:click|preventDefault={() => onDelete(i)}>Delete</a
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
