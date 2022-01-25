import { api } from './config'

export default {
    listarPorPaciente:(pacienteId) => {
        return api.get(`consultas/client/${pacienteId}`)
    },

    listarPorMedico:(medicoId) => {
        return api.get(`consultas/medic/${medicoId}`)
    }
}