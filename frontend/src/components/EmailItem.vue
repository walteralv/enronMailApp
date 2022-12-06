
<script setup>
import { computed, ref } from 'vue'

const maxBodyLength = 150
const isExpanded = ref(false)

const props = defineProps({
  email: Object, 
  term: String, 
})

const ContentPreview = computed(() => {
  return props.email.content.slice(0,maxBodyLength)
})

const displayContent = ref(ContentPreview.value)

const date = computed(() => {
  const date = new Date(props.email.date)
  return date.toLocaleDateString()
})

const expantToggle = () => {
  if (isExpanded.value){
    displayContent.value = ContentPreview.value
  }else{
    displayContent.value = props.email.content
  }
  isExpanded.value=!isExpanded.value

} 

</script>

<template>
  <li class="group">
    <div class="aspect-w-1 drop-shadow-md hover:drop-shadow-2xl  aspect-h-1 w-full my-10 p-10 rounded-lg text-black	 bg-gray-200 xl:aspect-w-7 xl:aspect-h-8">
      <div @click="expantToggle">
        <h4 class="text-left"> <b class="text-lg">From: </b> {{ email.from }} </h4>
        <h4 class="text-left"> <b class="text-lg">To: </b> {{ email.to }}</h4>
        <h4 class="text-left"> <b class="text-lg">Date: </b> {{ date }} </h4>
        <h3 class="font-bold text-xl">{{ email.subject }}</h3>
        <p class="text-center whitespace-pre-wrap	"> {{ displayContent }} </p>
      </div>
    </div>
  </li>
</template>



