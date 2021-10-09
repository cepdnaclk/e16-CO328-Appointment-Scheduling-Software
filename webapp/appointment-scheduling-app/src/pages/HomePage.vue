<template>
    <div className="main-container-home">
      <video
        src="../assets/Doctors.mp4"
        autoplay="true"
        loop="true"
        muted="true"
      />
      <h2 className="main-heading">Join the world's </h2>
      <p className="main-subheading">Appoiment sheduling journey ...</p>
      <div className="btn-container">
        <!-- <Link to="/hospitals"> -->
        <button className="btn btn-outline">FIND SERVICE</button>
        <!-- </Link> -->
        <!-- <Link to="/about"> -->
        <router-link to="/about">
        <button className="btn btn-solid">ABOUT US</button>
        </router-link>
        <!-- </Link> -->
      </div>
    </div>
</template>

<script>
export default {

}
</script>

<style>
@import url('../assets/homepage.css');
</style>