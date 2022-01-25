<template>
  <div id="pacienteVerMedicos">
    <NavbarPaciente/>
    <div class="barra-pesquisa">
      <form>
        <select v-model="filter.especialidade">
          <option v-for="(esp,index) in especialidades" :key="index">{{esp}}</option>
        </select>
        <select v-model="filter.estado">
          <option v-for="(estado,index) in estados" :key="index">{{estado}}</option>
        </select>
        <input type="text" placeholder="Cidade" v-model="filter.cidade"/>
      </form>
    </div>
  </div>
</template>


<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
import Medico from '../../services/medicos'

export default {
  data(){
    return{
      medicos: [],
      filter: {especialidade: '', estado: '', cidade: ''},
      especialidades: ['Ortopedia','Pediatria','Dermatologia','Geriatria']
    }
  },
  mounted(){
    Medico.listar().then(res => {
      console.log(res.data)
      this.medicos = res.data
    })
  },
  components:{
    NavbarPaciente
  }
}
</script>