<template>
  <nav class="sticky top-0 w-full bg-gradient-to-r from-green-900 to-blue-500 px-4 py-4 text-white flex">
    <div>
      <a href="" class=" text-black text-xl font-black rounded-lg border-2 px-2 border-solid border-black">AppOIMENTmE</a>
    </div>
      <div class=" px-9 py-2">
        <button  class=" px-3 hover:text-gray-500 focus:outline-none " @click="$emit('open-login')">LOGIN</button>
        <button  class=" px-3 hover:text-gray-500 focus:outline-none"  @click="$emit('open-signup')">SIGN UP</button>
      </div>
  </nav>
</template>

<script>
export default {

}
</script>

<style>

</style>