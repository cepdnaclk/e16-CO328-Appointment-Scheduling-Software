
<template>
      <section class="z-20 h-screen w-screen bg-gray-500 fixed top-0 opacity-70" @click="$emit(eventName)" />
      <div class=" absolute inset-0">
          <div class=" flex h-full">
              <div class="z-30 m-auto bg-blue-100 p-2 rounded shadow w-1/3">
                  <slot></slot>
              </div>
          </div>
      </div>
</template>

<script>
export default {
    props:{
        "eventName":String,
    }
}
</script>

<style>

</style>