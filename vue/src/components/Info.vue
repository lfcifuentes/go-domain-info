<template>
  <b-container>
    <b-card title="Listado de consultas" class="mb-2"  >
    </b-card>
    <b-card>
      <b-table striped bordered :small="true" :busy="loading"  hover :items="filtrerData" :fields="fields">
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
      loading:true,
      data:[],
      fields: {
        Url:{
          label: 'Url',
          sortable: true
        },
        Title:{
          label:'Titulo',
          sortable:true
        },
        IsDown:{
          label:'Estado',
          sortable:true
        }
      }
    }
  },
  computed:{
    filtrerData(){
      let arr = [];
      for (var $i=0;$i<this.data.length;$i++){
        arr.push({
          Url:this.data[$i].Url,
          Title:this.data[$i].Title,
          IsDown:this.data[$i].IsDown?'Caido':'En linea',
          _rowVariant: this.data[$i].IsDown?'danger':''
        })
      }
      return arr ;
    }
  },
  mounted(){
    this.loading = true;
    axios.post("http://localhost:3333/servers")
      .then((res)=>{
        this.loading = false;
        this.data = (res.data.Items);
      })
  }
}
</script>

<style scoped>

</style>
