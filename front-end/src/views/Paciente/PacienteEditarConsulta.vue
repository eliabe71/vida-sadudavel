<template>
  <div class="container">
    <form class="form">
       <div class="mb-3">
        <label for="medicoInput" class="form-label">Médico:</label>
        <input type="text" class="form-control" id="medicoInput" :value="$route.params.consulta.medico.nome" disabled/>
      </div>
      <div class="mb-3">
        <label for="dataInput" class="form-label">Data:</label>
        <input type="date" class="form-control" id="dataInput" :value="$route.params.consulta.data" />
      </div>
    </form>
  </div>
</template>