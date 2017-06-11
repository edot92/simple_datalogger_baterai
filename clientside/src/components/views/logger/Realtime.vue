<template lang="html">
<div id="realtime" class="content">
  <div class="row">
      <div class="col-md-4">
          <div class="form-group">
              <label class="control-label col-sm-5" for="pwd">Temperature:</label>
              <div class="col-sm-6">



                  <input type="text" class="form-control" v-model="Temperature" placeholder="???">
              </div>
          </div>
      </div>
      <div class="col-md-4">
          <div class="form-group">
              <label class="control-label col-sm-5" for="pwd">Humidity:</label>
              <div class="col-sm-6">
                  <input type="text" class="form-control"               v-model="Humidity"
 placeholder="???">
              </div>
          </div>
      </div>
      <div class="col-md-4">
          <div class="form-group">
              <label class="control-label col-sm-5" for="pwd"> Tegangan:</label>
              <div class="col-sm-6">
                  <input type="text" class="form-control"   v-model="Loadvoltage"  placeholder="???">
              </div>
          </div>
          </div>
      </div>
      <div class="row">
      <div class="col-md-4">
          <div class="form-group">
              <label class="control-label col-sm-5" for="pwd">Arus:</label>
              <div class="col-sm-6">
                  <input type="text" class="form-control" v-model="Current" placeholder="???">
              </div>
          </div>
      </div>
      <div class="col-md-4">
          <div class="form-group">
              <label class="control-label col-sm-5" for="pwd">Daya:</label>
              <div class="col-sm-6">
                  <input type="text" class="form-control"  v-model="Daya"  placeholder="???">
              </div>
          </div>
          </div>
      </div>
      <br>

<div class="row">
  <div class="col-md-6">
    <div class="form-group">
   <label class="control-label col-sm-5" for="email">Pilihan Grafik:</label>
   <div class="col-sm-6">
     <select class="form-control" v-model="dropdownRealtime">
  <option v-for="option in dropdownRealtimeOption" v-bind:value="option.value">
    {{ option.text }}
  </option>
</select>

     <!-- <select class="form-control" id="dropdownRealtime">
<option>Tegangan</option>
<option>Arus</option>
<option>Kelembapan</option>
</select> -->
   </div>
 </div>
  </div>
  <div class="col-md-3">
    <button type="button" class="btn btn-primary btn-block" @click.prevent="showChart()">
<i class="fa fa-sign-in">Tampilkan</i>
    </button>
  </div>

</div>
<br>
  <div class="row">
    <template>
<vue-highcharts :options="options" ref="lineCharts"></vue-highcharts>
<button @click="load">load</button>
</template>
  </div>
</div>
</template>

<script>
import VueHighcharts from 'vue2-highcharts'
// const asyncData = {
//   name: 'Tokyo',
//   marker: {
//     symbol: 'square'
//   },
//   data: [7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, {
//     y: 26.5,
//     marker: {
//       symbol: 'url(http://www.highcharts.com/demo/gfx/sun.png)'
//     }
//   }, 23.3, 18.3, 13.9, 9.6]
// }
export default{
  components: {
    VueHighcharts
  },
  mounted () {
    var thisVue = this

    setInterval(function () {
      thisVue.axios({
        method: 'get',
        url: '/device/realtimebox'
      })
    .then(function (response) {
      // console.log(response)
      thisVue.Temperature = response.data.data.temperature
      thisVue.Humidity = response.data.data.humidity
      if (response.data.data.current < 0) {
        response.data.data.current = response.data.data.current * -1
      }
      thisVue.Current = response.data.data.current
      thisVue.Shuntvoltage = response.data.data.shuntvoltage
      thisVue.Busvoltage = response.data.data.busvoltage
      thisVue.Loadvoltage = response.data.data.loadvoltage
      thisVue.Daya = parseFloat(thisVue.Current) * parseFloat(thisVue.Loadvoltage)
    })
    .catch(function (error) {
      console.error(error)
    })
    }, 2000)
  },
  data () {
    return {
      Daya: 0,
      Temperature: 4.42,
      Humidity: 4.42,
      Current: 1,
      Shuntvoltage: 4.42,
      Busvoltage: 4.42,
      Loadvoltage: 4.42,
      dropdownRealtime: 'Tegangan',
      dropdownRealtimeOption: [
        { text: 'Tegangan', value: 'Tegangan' },
        { text: 'Arus', value: 'Arus' },
        { text: 'Kelembapan', value: 'Kelembapan' }
      ],
      options: {
        chart: {
          type: 'spline'
        },
        title: {
          text: 'Tittle ??'
        },
        // subtitle: {
        //   text: 'Source: WorldClimate.com'
        // },
        xAxis: {
          categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun',
            'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
        },
        yAxis: {
          title: {
            text: 'Temperature'
          },
          labels: {
            formatter: function () {
              return this.value + '°'
            }
          }
        },
        tooltip: {
          crosshairs: true,
          shared: true
        },
        credits: {
          enabled: false
        },
        plotOptions: {
          spline: {
            marker: {
              radius: 4,
              lineColor: '#666666',
              lineWidth: 1
            }
          }
        },
        series: []
      }
    }
  },
  methods: {
    load () {
      // let lineCharts = this.$refs.lineCharts
      // lineCharts.delegateMethod('showLoading', 'Loading...')
      // setTimeout(() => {
      //   lineCharts.addSeries(asyncData)
      //   lineCharts.hideLoading()
      // }, 2000)
    },
    showChart () {
      const param = {
        jenis_chart: this.dropdownRealtime
      }
      var thisVue = this
      this.axios({
        method: 'post',
        url: '/device/chart',
        data: param
      })
    .then(function (response) {
      var unitYTime = []
      var valueX = []
      for (var i = 0; i < response.data.lenght; i++) {
        unitYTime[i] = response.data[i].time
        valueX[i] = response.data[i].value
      }
      // console.log(response.data)
      // var options = {
      //   title: {
      //     text: 'Monthly Average Temperature'
      //   },
      //   xAxis: {
      //     categories: unitYTime
      //   },
      //   yAxis: {
      //     title: {
      //       text: 'Temperature'
      //     },
      //     labels: {
      //       formatter: function () {
      //         return this.value + '°'
      //       }
      //     }
      //   }
      // }
      let lineCharts = thisVue.$refs.lineCharts
      lineCharts.delegateMethod('showLoading', 'Loading...')
      const dataChart = {
        name: thisVue.dropdownRealtime,
        marker: {
          symbol: 'square'
        },

        data: valueX
      }
      // setTimeout(() => {
      lineCharts.addSeries(dataChart)
      lineCharts.hideLoading()
      // }, 2000)
      // lineCharts.mergeOption(options)
      // console.log(response.data)
    })
    .catch(function (error) {
      console.log(error)
    })
    }
  }
}
</script>

<style lang="css">
</style>
