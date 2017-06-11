<template lang="html">
<div id="history" class="content">
  <form action="" method="#" class="form-horizontal" role="form">

      <div class="form-group">
          <label class="control-label col-sm-2" for="pwd">Range Waktu:</label>
          <div class="col-sm-10">
              <input type="text" class="form-control" id="demo" placeholder="Pilih Waktu">
          </div>
      </div>

      <div class="form-group">
          <div class="col-sm-10 col-sm-offset-2">
              <button @click.prevent="showHistoryByDate()" class="btn btn-primary btn-block">Tampilkan Data</button>
          </div>
      </div>
  </form>
  <div class="row">
      <div class="table-responsive">
        <div id="people">
          <v-client-table :data="tableData" :columns="columns" :options="options"></v-client-table>
        </div>
      </div>
  </div>
</div>
</template>

<script>
var paramDate = {
  startdate: '',
  enddate: ''
}
export default {
  mounted () {
    var thisVue = this
    thisVue.axios({
      method: 'get',
      url: '/device/historylast'
    })
  .then(function (response) {
    for (var i = 0; i < response.data.data.length; i++) {
      response.data.data[i].id = i + 1
      response.data.data[i].created_at = moment(moment(response.data.data[i].created_at).toString()).format('DD-MM-YYYY hh:mm:ss')
    }
    thisVue.tableData = response.data.data
  }).catch(function (error) {
    console.log(error)
  })
    $('#demo').daterangepicker({
      locale: {
        format: 'DD-MM-YYYY HH:mm:ss'
      },
      'dateLimit': {
        'minutes': 30 * 24 * 60
      },
      'timePicker': true,
      'timePicker24Hour': true,
      'timePickerSeconds': true,
      'startDate': '11-06-2017',
      'endDate': '12-06-2017'
    }, function (start, end, label) {
    // console.log('New date range selected: ' + start.format('YYYY-MM-DD hh:mm:ss') + ' to ' + end.format('YYYY-MM-DD hh:mm:ss') + ' (predefined range: ' + label + ')')
      paramDate.startdate = start.format('YYYY-MM-DD HH:mm:ss.SSS')
      paramDate.enddate = end.format('YYYY-MM-DD HH:mm:ss.SSS')
      // console.log(paramDate)
    })
  },
  data () {
    return {
      columns: [
        'id',
        'temperature',
        'humidity',
        'current',
        'shuntvoltage',
        'busvoltage',
        'loadvoltage',
        'created_at'
        // 'update_at'
      ],
      tableData: [

      ],
      options: {
            // see the options API
      }
    }
  },
  methods: {
    showHistoryByDate: function () {
      var thisVue = this
      thisVue.axios({
        method: 'post',
        url: '/device/historybydaterange',
        data: {
          startdate: paramDate.startdate,
          enddate: paramDate.enddate
        }

      })
    .then(function (response) {
      for (var i = 0; i < response.data.data.length; i++) {
        response.data.data[i].id = i + 1
        response.data.data[i].created_at = moment(moment(response.data.data[i].created_at).toString()).format('DD-MM-YYYY hh:mm:ss')
      }
      thisVue.tableData = response.data.data
    }).catch(function (error) {
      console.log(error)
    })
    }
  }
}
document.addEventListener('DOMContentLoaded', function () {

})
</script>

<style lang="css">
</style>
