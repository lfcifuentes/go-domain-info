<template>
  <b-container>
    <b-card title="Consultar Dominio" class="mb-2"  >
      <b-card-text>
        Ingrese el nombre del dominio que decea consultar.
      </b-card-text>
      <b-input-group prepend="Dominio" class="mt-3">
        <b-form-input v-model="url"></b-form-input>
        <b-input-group-append>
          <b-button variant="outline-success" v-on:click="send">
            Buscar
          </b-button>
        </b-input-group-append>
      </b-input-group>
    </b-card>
    <b-card v-show="show">
      <b-card class="text-left mb-2">
        <b-card-text>
          <div class="w-100 text-center" v-show="data.Logo != ''">
            <img v-bind:src="data.Logo" alt="" style="max-height:90px;margin: auto">
          </div>
          {{ data.Title }} <br>
          <b-badge variant="warning" v-show="data.IsDown">
            Servidor Caido
          </b-badge>
          <span v-if="data.ServersChange" >Cambio: {{ data.ServersChange}}</span>
          <span v-if="data.SslGrade != ''" >Grado Ssl: {{ data.SslGrade }}</span>
          <span v-if="data.PreviousSslGrade != ''">Grado Ssl anterior: {{ data.PreviousSslGrade }}</span>
        </b-card-text>
      </b-card>
      <b-table striped bordered :small="true" :busy="loading"  hover :items="data.endpoints" :fields="fields">
        <div slot="table-busy" class="text-center text-danger my-2">
          <b-spinner class="align-middle"></b-spinner>
          <strong>Loading...</strong>
        </div>
      </b-table>
    </b-card>
  </b-container>
</template>

<script>
  import axios from 'axios';
export default {
  data(){
    return {
      show:false,
      loading:false,
      url:"truora.com",
      data:{
        IsDown: false,
        Logo: "",
        PreviousSslGrade: "",
        ServersChange: false,
        SslGrade: "",
        Title: "",
        endpoints:[
        ]
      },
      fields: {
        ipAddress:{
          label: 'Dirección',
          sortable: true
        },
        country:{
          label:"País",
          sortable:true
        },
        grade:{
          label:"Certificado SSL:",
          sortable:true
        },
        owner:{
          label:"Propietario",
          sortable:true
        }
      }
    }
  },
  methods:{
    send(){
      this.show = true;
      this.loading = true;
      axios.post("http://localhost:3333/search",{
        "url": this.url
      }).then((res)=>{
        this.loading = false;
        this.data = (res.data);
      })
    }
  }
}
</script>

<style scoped>

</style>
