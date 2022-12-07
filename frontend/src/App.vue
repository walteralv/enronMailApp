<script setup>
import { defineComponent, ref } from 'vue'
import SearchBar from './components/SearchBar.vue'
import SearchResults from "./components/SearchResults.vue"
import EmailItem from "./components/EmailItem.vue"
import axios from 'axios'

const api = axios.create({
    baseURL: 'http://localhost:8000/'
})

const searchEmails = (term="") => {
    return api.get(`/emails/search?q=${term}`)
}

const emails = ref([])
const term = ref('')

const getEmails = async () => {
  const response = await searchEmails(term.value)
  emails.value = response.data.emails
  console.log(emails.value)
}

const searchNewTerm = (e) => {
  term.value = e.target.term.value
  getEmails()
}
</script>

<template>
  <div>
    <h1 class="text-green-400	text-6xl font-bold tracking-widest	">Email Search Engine</h1>
    <p class="text-left font-medium my-5">
      This application is a search engine that searches through the emails in the 
      <a class="" href="">Enron database</a>
      It contains data from about 150 users, mostly senior management of Enron, organized into folders. The corpus contains a total of about 0.5M messages. This data was originally made public, and posted to the web, by the Federal Energy Regulatory Commission during its investigation.
    </p>      
  </div>

  <SearchBar @submit.prevent="searchNewTerm" />
  <SearchResults :emails="emails" :term="term" :loading="false" />

</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
