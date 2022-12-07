<script setup>
import { ArrowUpOnSquareIcon } from '@heroicons/vue/24/solid'
</script>


<template>
  <div class="flex  justify-center py-5 px-4 sm:px-6 lg:px-8">
    <div class="w-full max-w-lx">
      <form class="grid grid-cols-8 gap-0" >
        <div class="col-span-7 h-full -space-y-px rounded-md shadow-sm">
            <input type="text" name="term" autocomplete="search" required="True" class="relative h-full block w-full appearance-none px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-green-500 focus:outline-none focus:ring-green-500 sm:text-sm" 
            placeholder="Search a word..." />
        </div>
        <div class="basis-1/4">
          <button type="submit" class="relative flex h-full bg-green-600 py-2 px-4 text-sm font-medium text-white hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-grenn-500 focus:ring-offset-2">
            <span class="inset-y-0 left-0 h-full ">
              <ArrowUpOnSquareIcon class="h-7" />
            </span>
          </button>
        </div>
      </form>
      
    </div>
  </div>
</template>
