<template>
      <section class=" h-screen w-screen bg-gray-500 fixed top-0 opacity-70" @click="$emit(eventName)" style=" z-index: 100;"/>
      <div class=" absolute inset-0">
          <div class=" flex h-full">
              <div class=" bg-transparent p-2 rounded shadow w-1/3" :class="customClases" 
                style="z-index:110">
                  <slot></slot>
              </div>
          </div>
      </div>
</template>

<script>
export default {
    props:{
        "eventName":String,
        "customClases":String,
    }
}
</script>

<style>

</style>