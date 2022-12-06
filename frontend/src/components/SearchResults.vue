<script setup>
import EmailItem from '../components/EmailItem.vue'

const props = defineProps({
  emails: Object, 
  term: String, 
  loading: Boolean
})

</script>

<template>
  <h1 v-if="!emails?.length && loading" class="center">Loading Emails..</h1>
  <h1 v-else-if="!emails?.length" class="center">No Emails found!</h1>
  <ul v-else>
    <EmailItem
      v-for="email in emails" 
      :email="email"
      :term="term"
    />
  </ul>
</template>

<style scoped>
.center {
  display: grid;
  place-items: center;
  padding: 4rem 0;
}
</style>
