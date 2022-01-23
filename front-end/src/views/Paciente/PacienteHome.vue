<template>
    <div id="pacienteHome">
      <NavbarPaciente/>

        
    </div>
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
export default {
  components:{
    NavbarPaciente
  }
}
</script>

<style>
*{
  font-family: Arial, Sans-serif, times;
}

#pacienteHome {
  width: 100vw;
  height: 100vh;
  background: linear-gradient(to bottom right,#fff, var(--primary));
}
</style>