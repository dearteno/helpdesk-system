<!-- src/routes/login/+page.svelte -->
<script lang="ts">
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  
  // Form state
  let email = '';
  let password = '';
  let loading = false;
  let errorMessage = '';
  
  onMount(() => {
    // Clear any error message when the component mounts
    errorMessage = '';
  });
  
  // Login function
  async function handleLogin() {
    if (!email || !password) {
      errorMessage = 'Please enter both email and password';
      return;
    }
    
    loading = true;
    errorMessage = '';
    
    try {
      // Call login API endpoint
      const response = await fetch('/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });
      
      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.message || 'Login failed');
      }
      
      // Redirect to dashboard on successful login
      goto('/dashboard');
      
    } catch (error) {
      console.error('Login error:', error);
      errorMessage = error.message || 'Failed to login. Please try again.';
    } finally {
      loading = false;
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
  <div class="max-w-md w-full space-y-8">
    <div>
      <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
        Sign in to your account
      </h2>
      <p class="mt-2 text-center text-sm text-gray-600">
        Access the Helpdesk Management System
      </p>
    </div>
    
    <form class="mt-8 space-y-6" on:submit|preventDefault={handleLogin}>
      <input type="hidden" name="remember" value="true" />
      <div class="rounded-md shadow-sm -space-y-px">
        <div>
          <label for="email-address" class="sr-only">Email address</label>
          <input 
            id="email-address" 
            name="email" 
            type="email" 
            autocomplete="email" 
            required 
            class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" 
            placeholder="Email address"
            bind:value={email}
          />
        </div>
        <div>
          <label for="password" class="sr-only">Password</label>
          <input 
            id="password" 
            name="password" 
            type="password" 
            autocomplete="current-password" 
            required 
            class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" 
            placeholder="Password"
            bind:value={password}
          />
        </div>
      </div>

      {#if errorMessage}
        <div class="text-red-500 text-sm mt-2">
          {errorMessage}
        </div>
      {/if}

      <div>
        <button 
          type="submit" 
          class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          disabled={loading}
        >
          {#if loading}
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
              <!-- Spinner icon -->
              <svg class="animate-spin h-5 w-5 text-indigo-300" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
            </span>
            Signing in...
          {:else}
            Sign in
          {/if}
        </button>
      </div>
      
      <div class="flex items-center justify-between">
        <div class="text-sm">
          <a href="/register" class="font-medium text-indigo-600 hover:text-indigo-500">
            Don't have an account? Register
          </a>
        </div>
      </div>
    </form>
  </div>
</div>
