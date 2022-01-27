import Login from '../views/Login.vue'
import PacienteHome from '../views/Paciente/PacienteHome.vue'
import PacienteEditarConsulta from '../views/Paciente/PacienteEditarConsulta.vue'
import PacienteHistorico from '../views/Paciente/PacienteHistorico.vue'
import PacienteVerMedicos from '../views/Paciente/PacienteVerMedicos.vue'
import MedicoHome from '../views/Medico/MedicoHome.vue'
import MedicoSolicitacoes from '../views/Medico/MedicoSolicitacoes.vue'
import PacienteAgendarConsulta from '../views/Paciente/PacienteAgendarConsulta.vue'

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
  },
  {
    path: '/paciente/agendarConsulta',
    name: 'pacienteAgendarConsulta',
    component: PacienteAgendarConsulta,
    params: true
  },
  {
    path: '/medicoHome',
    name: 'medicoHome',
    component: MedicoHome
  },
  {
    path: '/medico/solicitacoes',
    name: 'medicoSolicitcoes',
    component: MedicoSolicitacoes
  }
]