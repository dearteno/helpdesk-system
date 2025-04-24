<!-- src/routes/tickets/+page.svelte -->
<script lang="ts">
  import { onMount } from 'svelte';
  
  // Ticket types from UML
  const ticketTypes = ['Support', 'Bug', 'Feature Request', 'Question', 'Other'];
  
  // Ticket status options
  const statusOptions = ['Open', 'In Progress', 'Resolved', 'Closed'];
  
  // Pagination state
  let currentPage = 1;
  let totalPages = 1;
  let perPage = 10;
  
  // Table state
  let tickets = [];
  let loading = true;
  let error = null;
  
  // Filter state
  let statusFilter = '';
  let searchQuery = '';
  
  // New ticket form
  let showNewTicketModal = false;
  let newTicket = {
    type: ticketTypes[0],
    description: '',
  };
  
  // Fetch tickets data
  async function fetchTickets() {
    loading = true;
    error = null;
    
    try {
      const queryParams = new URLSearchParams({
        page: currentPage.toString(),
        per_page: perPage.toString(),
      });
      
      if (statusFilter) {
        queryParams.append('status', statusFilter);
      }
      
      if (searchQuery) {
        queryParams.append('search', searchQuery);
      }
      
      const response = await fetch(`/api/v1/ticket?${queryParams.toString()}`);
      
      if (!response.ok) {
        throw new Error('Failed to fetch tickets');
      }
      
      const data = await response.json();
      tickets = data.tickets;
      
      // Update pagination
      totalPages = data.pagination.total_pages;
    } catch (err) {
      console.error('Error fetching tickets:', err);
      error = err.message;
      tickets = [];
    } finally {
      loading = false;
    }
  }
  
  // Create new ticket
  async function createTicket() {
    try {
      const response = await fetch('/api/v1/ticket', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(newTicket),
      });
      
      if (!response.ok) {
        throw new Error('Failed to create ticket');
      }
      
      // Reset form and close modal
      newTicket = {
        type: ticketTypes[0],
        description: '',
      };
      showNewTicketModal = false;
      
      // Refresh tickets
      await fetchTickets();
      
    } catch (err) {
      console.error('Error creating ticket:', err);
      alert(`Error creating ticket: ${err.message}`);
    }
  }
  
  // Handle pagination
  function goToPage(page) {
    if (page < 1 || page > totalPages) return;
    currentPage = page;
    fetchTickets();
  }
  
  // Apply filters
  function applyFilters() {
    currentPage = 1; // Reset to first page when filtering
    fetchTickets();
  }
  
  // Format date for display
  function formatDate(dateString) {
    return new Date(dateString).toLocaleString();
  }
  
  // Lifecycle
  onMount(() => {
    fetchTickets();
  });
</script>

<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
  <div class="header py-5">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-gray-900">Tickets</h1>
      <button 
        class="px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700"
        on:click={() => showNewTicketModal = true}
      >
        Create New Ticket
      </button>
    </div>
    
    <!-- Filters -->
    <div class="mt-4 flex flex-col sm:flex-row gap-4">
      <div class="flex-1">
        <label for="search" class="sr-only">Search</label>
        <input 
          type="text" 
          id="search" 
          placeholder="Search tickets..." 
          bind:value={searchQuery}
          class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
        />
      </div>
      
      <div class="w-full sm:w-48">
        <label for="status" class="sr-only">Filter by status</label>
        <select 
          id="status" 
          bind:value={statusFilter}
          class="shadow-sm focus:ring-indigo-500 focus:border-indigo-500 block w-full sm:text-sm border-gray-300 rounded-md"
        >
          <option value="">All Statuses</option>
          {#each statusOptions as status}
            <option value={status.toLowerCase()}>{status}</option>
          {/each}
        </select>
      </div>
      
      <div>
        <button 
          class="w-full px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50"
          on:click={applyFilters}
        >
          Apply Filters
        </button>
      </div>
    </div>
  </div>
  
  {#if loading}
    <div class="flex justify-center py-12">
      <svg class="animate-spin h-8 w-8 text-indigo-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
    </div>
  {:else if error}
    <div class="bg-red-50 border-l-4 border-red-400 p-4 my-4">
      <div class="flex">
        <div class="flex-shrink-0">
          <svg class="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
            <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm1-11a1 1 0 10-2 0v4a1 1 0 102 0V7zm-1 8a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd" />
          </svg>
        </div>
        <div class="ml-3">
          <p class="text-sm text-red-700">
            {error}
          </p>
        </div>
      </div>
    </div>
  {:else if tickets.length === 0}
    <div class="py-12 text-center">
      <p class="text-gray-500">No tickets found. Create a new ticket to get started.</p>
    </div>
  {:else}
    <!-- Tickets Table -->
    <div class="mt-8 flex flex-col">
      <div class="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
          <div class="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
            <table class="min-w-full divide-y divide-gray-300">
              <thead class="bg-gray-50">
                <tr>
                  <th scope="col" class="py-3 pl-4 pr-3 text-left text-xs font-medium uppercase tracking-wide text-gray-500 sm:pl-6">ID</th>
                  <th scope="col" class="px-3 py-3 text-left text-xs font-medium uppercase tracking-wide text-gray-500">Type</th>
                  <th scope="col" class="px-3 py-3 text-left text-xs font-medium uppercase tracking-wide text-gray-500">Description</th>
                  <th scope="col" class="px-3 py-3 text-left text-xs font-medium uppercase tracking-wide text-gray-500">Status</th>
                  <th scope="col" class="px-3 py-3 text-left text-xs font-medium uppercase tracking-wide text-gray-500">Created</th>
                  <th scope="col" class="relative py-3 pl-3 pr-4 sm:pr-6">
                    <span class="sr-only">Actions</span>
                  </th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-200 bg-white">
                {#each tickets as ticket}
                  <tr>
                    <td class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-6">{ticket.id.substring(0, 8)}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{ticket.type}</td>
                    <td class="px-3 py-4 text-sm text-gray-500 max-w-xs truncate">{ticket.description}</td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm">
                      <span class={`inline-flex rounded-full px-2 text-xs font-semibold leading-5 ${
                        ticket.status === 'open' ? 'bg-green-100 text-green-800' :
                        ticket.status === 'in progress' ? 'bg-yellow-100 text-yellow-800' :
                        ticket.status === 'resolved' ? 'bg-blue-100 text-blue-800' :
                        'bg-gray-100 text-gray-800'
                      }`}>
                        {ticket.status}
                      </span>
                    </td>
                    <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">{formatDate(ticket.created_at)}</td>
                    <td class="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                      <a href={`/tickets/${ticket.id}`} class="text-indigo-600 hover:text-indigo-900">View</a>
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Pagination -->
    <div class="flex items-center justify-between border-t border-gray-200 bg-white px-4 py-3 sm:px-6 mt-4">
      <div class="flex flex-1 justify-between sm:hidden">
        <button 
          on:click={() => goToPage(currentPage - 1)}
          disabled={currentPage === 1}
          class="relative inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
          class:opacity-50={currentPage === 1}
        >
          Previous
        </button>
        <button 
          on:click={() => goToPage(currentPage + 1)}
          disabled={currentPage === totalPages}
          class="relative ml-3 inline-flex items-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
          class:opacity-50={currentPage === totalPages}
        >
          Next
        </button>
      </div>
      <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
        <div>
          <p class="text-sm text-gray-700">
            Showing <span class="font-medium">{(currentPage - 1) * perPage + 1}</span> to <span class="font-medium">{Math.min(currentPage * perPage, tickets.length)}</span> of{' '}
            <span class="font-medium">{tickets.length}</span> results
          </p>
        </div>
        <div>
          <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
            <button
              on:click={() => goToPage(currentPage - 1)}
              disabled={currentPage === 1}
              class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              class:opacity-50={currentPage === 1}
            >
              <span class="sr-only">Previous</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M12.79 5.23a.75.75 0 01-.02 1.06L8.832 10l3.938 3.71a.75.75 0 11-1.04 1.08l-4.5-4.25a.75.75 0 010-1.08l4.5-4.25a.75.75 0 011.06.02z" clip-rule="evenodd" />
              </svg>
            </button>
            
            {#each Array(totalPages).fill().map((_, i) => i + 1) as page}
              {#if page === 1 || page === totalPages || (page >= currentPage - 1 && page <= currentPage + 1)}
                <button
                  on:click={() => goToPage(page)}
                  class={`relative inline-flex items-center px-4 py-2 text-sm font-semibold ${
                    page === currentPage 
                      ? 'z-10 bg-indigo-600 text-white focus-visible:outline-indigo-600' 
                      : 'text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:outline-offset-0'
                  }`}
                >
                  {page}
                </button>
              {:else if page === currentPage - 2 || page === currentPage + 2}
                <span class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-700 ring-1 ring-inset ring-gray-300">
                  ...
                </span>
              {/if}
            {/each}
            
            <button
              on:click={() => goToPage(currentPage + 1)}
              disabled={currentPage === totalPages}
              class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
              class:opacity-50={currentPage === totalPages}
            >
              <span class="sr-only">Next</span>
              <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                <path fill-rule="evenodd" d="M7.21 14.77a.75.75 0 01.02-1.06L11.168 10 7.23 6.29a.75.75 0 111.04-1.08l4.5 4.25a.75.75 0 010 1.08l-4.5 4.25a.75.75 0 01-1.06-.02z" clip-rule="evenodd" />
              </svg>
            </button>
          </nav>
        </div>
      </div>
    </div>
  {/if}
  
  <!-- New Ticket Modal -->
  {#if showNewTicketModal}
    <div class="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center z-10">
      <div class="bg-white rounded-lg overflow-hidden shadow-xl transform transition-all sm:max-w-lg sm:w-full">
        <div class="px-4 pt-5 pb-4 sm:p-6">
          <div class="mt-3 text-center sm:mt-0 sm:text-left">
            <h3 class="text-lg leading-6 font-medium text-gray-900">Create New Ticket</h3>
            <div class="mt-4">
              <form on:submit|preventDefault={createTicket} class="space-y-4">
                <div>
                  <label for="type" class="block text-sm font-medium text-gray-700">Ticket Type</label>
                  <select 
                    id="type" 
                    bind:value={newTicket.type}
                    class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md"
                  >
                    {#each ticketTypes as type}
                      <option value={type}>{type}</option>
                    {/each}
                  </select>
                </div>
                
                <div>
                  <label for="description" class="block text-sm font-medium text-gray-700">Description</label>
                  <textarea 
                    id="description" 
                    rows="4"
                    bind:value={newTicket.description}
                    class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                    placeholder="Describe the issue..."
                    required
                  ></textarea>
                </div>
                
                <div class="flex justify-end space-x-3">
                  <button 
                    type="button"
                    class="bg-white py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                    on:click={() => showNewTicketModal = false}
                  >
                    Cancel
                  </button>
                  <button 
                    type="submit"
                    class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  >
                    Create Ticket
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  {/if}
</div>
