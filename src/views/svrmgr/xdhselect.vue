<template>
  <div>
    <el-select v-model="svalue" placeholder="游戏" >
      <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
        <span style="float: left">{{ item.label }}</span>
        <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
      </el-option>
    </el-select>
  </div>
</template>

<script>
/* eslint-disable */
export default {
    name: 'XdhSelect',
    data() {
        return {
            options:[
                {
                    value: '20201',
                    label: '方块部落'
                }, 
                {
                    value: '20301',
                    label: '国王'
                },
                {
                    value: '20401',
                    label: '春节红包'
                },
                {
                    value: '20501',
                    label: '休闲射击'
                },
                {
                    value: '20601',
                    label: '猫咪'
                },
                {
                    value: '20701',
                    label: '消灭星星'
                },
                {
                    value: '20801',
                    label: '挂机奇迹'
                },
                {
                    value: '20901',
                    label: '方块2'
                }
            ],
            svalue: ''
        }
    },
    methods: {
        
    },
    watch:{
        //判断下拉框的值是否有改变
        svalue(val, oldVal) {
            console.log('new: %s, old: %s', val, oldVal)
            if(val!=oldVal){
                this.$emit('input', this.svalue); 
            }
        }, 
    },
    props: {
        value: {
            type: String
        },//接受外部v-model传入的值
    },
    mounted(){
        //初始话下拉框的值
        this.svalue=this.value;
    }
}
</script>