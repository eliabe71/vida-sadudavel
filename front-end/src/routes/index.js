import Login from '../views/Login.vue'
import PacienteHome from '../views/Paciente/PacienteHome.vue'
import PacienteEditarConsulta from '../views/Paciente/PacienteEditarConsulta.vue'
import PacienteHistorico from '../views/Paciente/PacienteHistorico.vue'
import PacienteVerMedicos from '../views/Paciente/PacienteVerMedicos.vue'

export default [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/pacienteHome',
    name: 'PacienteHome',
    component: PacienteHome
  },
  {
    path: '/paciente/minhasConsultas/edit/:id',
    name: 'pacienteEditarConsulta',
    component: PacienteEditarConsulta,
    params: true
  },
  {
    path: '/paciente/historicoConsultas',
    name: 'pacienteHistoricoConsultas',
    component: PacienteHistorico
  },
  {
    path: '/paciente/verMedicos',
    name: 'pacienteVerMedicos',
    component: PacienteVerMedicos
  }
]