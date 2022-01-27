<template>
  <div id="medicoHome">
    <NavbarMedico/>
    <h2 class="title">Minhas Consultas</h2>
    <div class="bloco-consultas">
      <table>
        <tr v-for="(cons,index) in divide(consultas)" :key="index">
          <td v-for="c in cons" :key="c.id">
            <div class="card">
              <img src="../../assets/images/logo.png" class="card-img-top" alt="Logo"> 
              <div class="card-body">
                <h5 class="card-title">Consulta</h5>
                <p class="card-text"><b>Cliente:</b> {{c.clientName}}</p>
                <p class="card-text"><b>Data:</b> {{formatDate(c.day)}}</p>
                <p class="card-text"><b>Horário:</b> {{formatHour(c.hourInit)}}</p>
              </div>
            </div>
          </td>
        </tr>
      </table>
    </div>

  </div>
</template>

<script>
import NavbarMedico from '../../components/NavbarMedico/NavbarMedico.vue'
import Medicos from '../../services/medicos'
import Consultas from '../../services/consultas'

export default {
  data(){
    return{
      consultas: [],
      medicoLogado: {}
    }
  },
  mounted(){
    Medicos.getMedico(1).then(res => {
      this.medicoLogado = res.data.Data[0]
    }),
    Consultas.listarPorMedico(1).then(res => {
      this.consultas = res.data.Data.filter(c => {
        return (c.status === true && c.effected === false)
      })
    })
  },
  methods:{
    divide(cons){
      let res = []
      let aux = []
      for(let i=0; i<cons.length; i++){
        aux.push(cons[i])
        if(((i+1)%4 === 0 && i !== 0) || i === cons.length-1){
          res.push(aux)
          aux = []
        }
      }
      return res
    },
  
    formatDate(d){
      let res = d.split('T')
      res = res[0].split('-')
      res = res[2]+'/'+res[1]+'/'+res[0]
      return res
    },
    formatHour(hour){
      let res = hour.split('T')
      res = res[1].split(':')
      res = res[0]+':'+res[1]
      return res
    },
  },
  components: {
      NavbarMedico
  }
}

</script>

<style>
*{
  font-family: Arial, Sans-serif, times;
}

#medicoHome {
  width: 100vw;
  height: 100vh;
  background-color:  var(--primary);
}

.title{
  color: #fff;
  margin-top: 10px;
  margin-left: 20px;
  margin-bottom: 0;
  font-size: 25pt;
  font-weight: bold;
}

.bloco-consultas{
  display: flex;
}

.bloco-consultas table{
  margin-left: 10px;
  margin-right: 10px;
}

.bloco-consultas td{
  min-width: 18vw;
  padding: 10px;
}

.card{
  background-color: var(--background-default);
  display: flex;
  width: 100%;
  height: 100%;
}

.card img{
  width: 100%;
  height: 100%;
  margin-bottom: 0vw;
}

.card-body{
  margin-top: 0;
  width: 100%;
  height: 100%;
}

.card-title{
  font-size: 2vw;
  font-weight: bold;
  margin-bottom: 0.1vw;
  margin-top: 0;
}
.card-text{
  font-size: 1.5vw;
  margin-bottom: 0.05vw;
  margin-top: 0;
}

.card-footer{
  width: 100%;
  height: 3vw;
  display: flex;
  align-items: center;
}

</style>
