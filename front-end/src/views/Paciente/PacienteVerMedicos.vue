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
      </form>
      
      <div>
        <h2>Médicos</h2>
        <ul class="lista">
          <li class="item" v-for="med in filtraMedicos()" :key="med.id">
            <span class="atributo"><b>Nome</b>: {{med.name}} {{med.lastName}}</span>
            <span class="atributo"><b>Especialidade:</b> {{med.areaOfOcupation}}</span>
            <span class="atributo"><b>Preço por consulta:</b> {{med.price}}</span>
            <router-link :to="{name: 'pacienteAgendarConsulta',  params:{medico: med} }" >Agendar Consulta</router-link>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>


<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
import Medico from '../../services/medicos'
import Pacientes from '../../services/pacientes'

export default {
  data(){
    return{
      medicos: [],
      pacienteLogado: {},
      filter: {
        especialidade: null,
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
    }),
    Pacientes.getPaciente(1).then(res => {
      this.pacienteLogado = res.data.Data[0]
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
          return false
      }).filter(e => {
        if(this.filter.estado !== null && this.filter.estado !== undefined)
          return ( e.state.toLowerCase() === this.filter.estado.toLowerCase() )
        else
          return false
      }).filter(e => {
        if(this.filter.cidade !== null && this.filter.cidade !== undefined && this.filter.cidade !== '' )
          return ( e.city.toLowerCase().includes(this.filter.cidade.toLowerCase()) )
        else
          return false
      })
      return res
    }
  }
}
</script>

<style>

*{
  font-family: Arial, Sans-serif, times;
}

#pacienteVerMedicos {
  width: 100vw;
  height: 100vh;
  background-color:  var(--primary);
}

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