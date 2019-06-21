<template>
    <div class="table">
        <div class="app-container">
            <div class="handle-box">
                <el-row :gutter="10">
                    <el-col :span="4">
                        <xdh-select v-model="selected_game"> </xdh-select>
                    </el-col>
                    <el-col :span="2">
                        <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                    </el-col>
                </el-row>
            </div>
            <div class="handle-box">
                <el-row :gutter="10">
                    <el-col :span="4">
                        <el-input placeholder="请输入生成数量" v-model="input_number"> </el-input>
                    </el-col>
                    <el-col :span="4">
                        <el-input placeholder="请输入奖励类型" v-model="input_atype"> </el-input>
                    </el-col>
                    <el-col :span="4">
                        <el-input placeholder="请输入奖励数量" v-model="input_acount"> </el-input>
                    </el-col>
                    <el-col :span="2" >
                        <el-button type="primary"  @click="handleGen">生成<i class="el-icon-upload el-icon--right"></i></el-button>
                    </el-col>
                </el-row>
            </div>

            <el-table :data="data" border class="table">
                <el-table-column prop="game" label="游戏" width="100">
                </el-table-column>
                <el-table-column prop="code" label="礼包码" width="160">
                </el-table-column>
                <el-table-column prop="atype" label="礼包类型" width="200">
                </el-table-column>
                <el-table-column prop="acount" label="奖励数量" width="300">
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="50">
                </el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
/* eslint-disable */
import XdhSelect from './xdhselect'
export default {
    components: { XdhSelect },
    name: 'gifts',
    data() {
        return {
            tableData: [],
            cur_page: 1,
            is_search: false,
            selected_game : '',
            actioning : 'get',
            input_number : 1,
            input_atype : '1',
            input_acount : 1,
            form: {
                game: this.selected_game,
                number : 1,//this.input_number,
                atype: '1', //this.input_atype,
                acount: 1, //this.input_acount
            },
            idx: -1
        }
    },
    created() {

    },
    computed: {
        data() {
            return this.tableData;
        }
    },
    mounted(){

    },
    methods: {
        // 分页导航
        handleCurrentChange(val) {
            this.cur_page = val;
            this.getData();
        },
        getData() {
            this.handleSearch()
        },

        handleSearch() {
            this.is_search = true;
            const user = this.$store.getters.name
            const page = this.$route.path
            const game = this.selected_game
            console.log(this.selected_game)
            this.tableData = []
            this.$store.dispatch('svrmgr/Getgifts', { user, page, game}).then((res) => {
                console.log(res)
                if (res.code != 0) {
                    this.$message.error('服务器返回失败:' + res.errorCode)
                } else {
                    if (res.list) {
                        console.log( this.tableData)
                        this.tableData = res.list
                    } else {
                        this.$message.error('没有数据')
                    }
                } 
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });
        },
        handleGen() {
            if (this.selected_game === '') {
                this.$message.success(`请先选择游戏`)
                return
            }
            this.form = {
                game: this.selected_game,
                number : this.input_number,
                atype: this.input_atype,
                acount: this.input_acount
            }
            this.actioning = 'gen'
            this.saveEdit()
        },
        // 保存编辑
        saveEdit() {
            const user = this.$store.getters.name
            const formdata = {
                game : this.selected_game,
                number : this.form.number,
                atype : this.form.atype,
                acount : this.form.acount,
                action : this.actioning
            }
            console.log(formdata)
            this.$store.dispatch('svrmgr/Gengifts', { user, formdata}).then((res) => {
                if (res.errorCode != 0) {
                    this.$message.error(`执行失败 :` + res.errorCode);
                } else {
                    this.$message.success(`执行成功`);
                }
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });

            this.handleSearch()
        },
    }
}
</script>

<style scoped>
    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 120px;
    }

    .handle-input {
        width: 300px;
        display: inline-block;
    }
    .del-dialog-cnt{
        font-size: 16px;
        text-align: center
    }
    .table{
        width: 100%;
        font-size: 14px;
    }
    .red{
        color: #ff0000;
    }
    .mr10{
        margin-right: 10px;
    }
</style>
