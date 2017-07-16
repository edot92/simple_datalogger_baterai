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
    <download-excel
    class   = "btn btn-primary"
    :data   = "json_data"
    :fields = "json_fields"
    name    = "filename.xls">
    Download Excel
</download-excel>
  </div>
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
import JsonExcel from 'vue-json-excel'
var paramDate = {
  startdate: '',
  enddate: ''
}
export default {
  components: {
    'downloadExcel': JsonExcel
  },
  mounted () {
    var thisVue = this
    thisVue.axios({
      method: 'get',
      url: '/device/historylast'
    }).then(function (response) {
      for (var i = 0; i < response.data.data.length; i++) {
        response.data.data[i].id = i + 1
        response.data.data[i].created_at = moment(moment(response.data.data[i].created_at).toString()).format('DD-MM-YYYY HH:mm:ss')
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
      paramDate.startdate = start.format('YYYY-MM-DD HH:mm:ss.SSS')
      paramDate.enddate = end.format('YYYY-MM-DD HH:mm:ss.SSS')
    })
  },
  data () {
    return {
      json_fields: {
        'id': 'String',
        'temperature': 'String',
        'humidity': 'String',
        'current': 'String',
        // 'shuntvoltage':'String',
        // 'busvoltage':'String',
        'loadvoltage': 'String',
        'created_at': 'String'
      },
      json_data: [
        {
          'id': 8903,
          'temperature': 28.5,
          'humidity': 90.4,
          'current': -0.3,
          // 'shuntvoltage': -0.02,
          // 'busvoltage': 0.98,
          'loadvoltage': 0.98,
          'created_at': '2017-07-16 00:18:08.000000000+07:00',
          'update_at': '2017-07-16 00:18:08.000000000+07:00'
        }
        // {
        //   'ID': 8904,
        //   'temperature': 28.5,
        //   'humidity': 90.4,
        //   'current': -0.1,
        //   'shuntvoltage': 0.01,
        //   'busvoltage': 0.92,
        //   'loadvoltage': 0.92,
        //   'created_at': '2017-07-16 00:18:10.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:10.000000000+07:00'
        // },
        // {
        //   'ID': 8905,
        //   'temperature': 28.5,
        //   'humidity': 90.3,
        //   'current': -0.2,
        //   'shuntvoltage': -0.02,
        //   'busvoltage': 0.99,
        //   'loadvoltage': 0.99,
        //   'created_at': '2017-07-16 00:18:13.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:13.000000000+07:00'
        // },
        // {
        //   'ID': 8906,
        //   'temperature': 28.5,
        //   'humidity': 90.3,
        //   'current': -0.2,
        //   'shuntvoltage': 0.02,
        //   'busvoltage': 0.92,
        //   'loadvoltage': 0.92,
        //   'created_at': '2017-07-16 00:18:15.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:15.000000000+07:00'
        // },
        // {
        //   'ID': 8907,
        //   'temperature': 28.5,
        //   'humidity': 90.4,
        //   'current': -0.2,
        //   'shuntvoltage': -0.01,
        //   'busvoltage': 0.98,
        //   'loadvoltage': 0.98,
        //   'created_at': '2017-07-16 00:18:17.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:17.000000000+07:00'
        // },
        // {
        //   'ID': 8908,
        //   'temperature': 28.5,
        //   'humidity': 90.4,
        //   'current': -0.2,
        //   'shuntvoltage': 0,
        //   'busvoltage': 0.98,
        //   'loadvoltage': 0.98,
        //   'created_at': '2017-07-16 00:18:20.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:20.000000000+07:00'
        // },
        // {
        //   'ID': 8909,
        //   'temperature': 28.5,
        //   'humidity': 90.4,
        //   'current': 0.1,
        //   'shuntvoltage': -0.02,
        //   'busvoltage': 0.98,
        //   'loadvoltage': 0.98,
        //   'created_at': '2017-07-16 00:18:22.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:22.000000000+07:00'
        // },
        // {
        //   'ID': 8910,
        //   'temperature': 28.5,
        //   'humidity': 90.4,
        //   'current': -0.5,
        //   'shuntvoltage': 0,
        //   'busvoltage': 0.94,
        //   'loadvoltage': 0.94,
        //   'created_at': '2017-07-16 00:18:24.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:24.000000000+07:00'
        // },
        // {
        //   'ID': 8911,
        //   'temperature': 28.5,
        //   'humidity': 90.5,
        //   'current': -0.2,
        //   'shuntvoltage': -0.05,
        //   'busvoltage': 0.92,
        //   'loadvoltage': 0.92,
        //   'created_at': '2017-07-16 00:18:27.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:27.000000000+07:00'
        // },
        // {
        //   'ID': 8912,
        //   'temperature': 28.5,
        //   'humidity': 90.5,
        //   'current': -0.3,
        //   'shuntvoltage': -0.02,
        //   'busvoltage': 1,
        //   'loadvoltage': 1,
        //   'created_at': '2017-07-16 00:18:29.000000000+07:00',
        //   'update_at': '2017-07-16 00:18:29.000000000+07:00'
        // }
      ],
      columns: [
        'id',
        'temperature',
        'humidity',
        'current',
        // 'shuntvoltage',
        // 'busvoltage',
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
          // thisVue.tableData =[]
          for (var i = 0; i < response.data.data.length; i++) {
            response.data.data[i].id = i + 1
            response.data.data[i].created_at = moment(moment(response.data.data[i].created_at).toString()).format('DD-MM-YYYY HH:mm:ss')
            // response.data.data[i].loadvoltage = null
            // response.data.data[i].shuntvoltage = null
            // json csv
            // thisVue.json_data[i].busvoltage = response.data.data[i].busvoltage
            // thisVue.json_data[i].created_at = response.data.data[i].created_at
            // thisVue.json_data[i].current = response.data.data[i].current
            // thisVue.json_data[i].humidity = response.data.data[i].humidity
            // thisVue.json_data[i].id = response.data.data[i].id
            // thisVue.json_data[i].loadvoltage = response.data.data[i].loadvoltage
            // // thisVue.json_data[i].shuntvoltage = response.data.data[i].shuntvoltage
            // thisVue.json_data[i].temperature = response.data.data[i].temperature
                        // response.data.data[i].ID = null
          }
          thisVue.tableData = response.data.data
          thisVue.json_data = thisVue.tableData
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
