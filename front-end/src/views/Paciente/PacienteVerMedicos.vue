<template>
  <div id="pacienteVerMedicos">
    <NavbarPaciente/>
    <div class="barra-pesquisa">
      <form>
        <select v-model="filter.especialidade" >
          <option :value="null">Selecione a especialidade</option>
          <option v-for="(esp,index) in especialidades" :key="index" :value="esp">{{esp}}</option>
        </select>
        <select v-model="filter.estado">
          <option v-for="(estado,index) in estados" :key="index">{{estado}}</option>
        </select>
        <input type="text" placeholder="Cidade" v-model="filter.cidade"/>
        <input type="time" v-model="medicos[0].hourinit">
      </form>
      
      <div>
        <h2>Médicos</h2>
        <ul class="lista">
          <li class="item" v-for="medico in filtraMedicos()" :key="medico.id">
            <span class="atributo"><b>Nome</b>: {{medico.name}} {{medico.lastName}}</span>
            <span class="atributo"><b>Especialidade:</b> {{medico.areaOfOcupation}}</span>
            <span class="atributo"><b>Preço por consulta:</b> {{medico.price}}</span>
            <button class="btn-success">Agendar Consulta</button>
          </li>
        </ul>
      </div>
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

      filter: {
        especialidade:'Ortopedia',
        estado: null,
        cidade: null
        },
      especialidades: ['Ortopedia','Pediatria','Dermatologia','Geriatria'],
      estados: ['CE','SP','AM']
    }
  },
  mounted(){
    Medico.listar().then(res => {
      this.medicos = res.data.Data
    })
  },
  components:{
    NavbarPaciente
  },
  methods:{
    filtraMedicos(){
      let res =  this.medicos.filter(e => {
        if(this.filter.especialidade !== null && this.filter.especialidade !== undefined)
          return ( e.areaOfOcupation.toLowerCase() === this.filter.especialidade.toLowerCase() )
        else
          return true
      }).filter(e => {
        if(this.filter.estado !== null && this.filter.estado !== undefined)
          return ( e.state.toLowerCase() === this.filter.estado.toLowerCase() )
        else
          return true
      }).filter(e => {
        if(this.filter.cidade !== null && this.filter.cidade !== undefined && this.filter.cidade !== '' )
          return ( e.city.toLowerCase().includes(this.filter.cidade.toLowerCase()) )
        else
          return true
      })
      return res
    }
  }
}
</script>

<style>

h2{
  margin-top: 20px;
  margin-left: 30px;
}

.lista{
  margin-top: 20px;
}

.item{
  display: flex;
  justify-content: space-around;
  align-items: center;
  border: 2px solid #47b132;
  padding: 10px;
  list-style: none;
  margin: 10px 30px 0px 0px;
}

.atributo{
  margin-right: 30px;
}

.btn-success{
  width: 200px;
}

</style>