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
                    <el-col :span="2" :offset="4">
                        <el-button type="primary"  @click="handleAdd">新增<i class="el-icon-upload el-icon--right"></i></el-button>
                    </el-col>
                </el-row>
            </div>
            <el-divider></el-divider>
            <el-divider></el-divider>
            <el-table :data="data" border class="table">
                <el-table-column prop="game" label="游戏" width="100">
                </el-table-column>
                <el-table-column prop="funcname" label="名字" width="100">
                </el-table-column>
                <el-table-column prop="funcswitch" label="限制次数" width="100">
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="10">
                </el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="50%">
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="名字">
                    <el-select v-model="form.funcname" placeholder="请选择" >
                        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                            <span style="float: left">{{ item.label }}</span>
                            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="form.funcswitch" placeholder="请选择" >
                        <el-option v-for="item in funcstatus" :key="item.value" :label="item.label" :value="item.value">
                            <span style="float: left">{{ item.label }}</span>
                            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                        </el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
/* eslint-disable */
import XdhSelect from './xdhselect'
export default {
    components: { XdhSelect },
    name: 'settings',
    data() {
        return {
            options:[
                {
                    value: 'maxPumpTimes',
                    label: 'maxPumpTimes'
                }
            ],
            funcstatus:[
                {
                    value: '5',
                    label: '5'
                },
                {
                    value: '4',
                    label: '4'
                },
                {
                    value: '3',
                    label: '3'
                },
                {
                    value: '2',
                    label: '2'
                },
                {
                    value: '1',
                    label: '1'
                },
                {
                    value: '0',
                    label: '0'
                }
            ],
            tableData: [],
            cur_page: 1,
            is_search: false,
            selected_game : '',
            editVisible: false,
            form: {
                game: this.selected_game,
                funcname : 'maxPumpTimes',
                funcswitch: '3'
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
            this.$store.dispatch('svrmgr/GetSvrSwitch', { user, page, game}).then((res) => {
                console.log(res)
                if (res.errorCode != 0) {
                    this.$message.error('服务器返回失败:' + res.errorCode)
                } else {
                    this.tableData = []
                    // this.tableData = res.data
                    for(var i = 0; i < res.data.length; i++) {
                        if (res.data[i].funcname == 'maxPumpTimes') {
                            let item = res.data[i]
                            item['game'] = game
                            this.tableData.push(item)
                        }
                    } 
                    // this.tableData = res.data.filter(egg => {
                    //     if (egg.funcname === 'maxPumpTimes') {
                    //         return true
                    //     }
                    //     return false
                    // })
                } 
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });
        },
        handleAdd() {
            if (this.selected_game === '') {
                this.$message.success(`请先选择游戏`)
                return
            }
            this.form = {
                game: this.selected_game,
                funcname : 'maxPumpTimes',
                funcswitch: '3'
            }
            this.editVisible = true
        },
        handleEdit(index, row) {
            this.idx = index;
            const item = this.tableData[index];
            this.form = {
                game: this.selected_game,
                funcname : item.funcname,
                funcswitch: item.funcswitch
            }
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
            const user = this.$store.getters.name
            const formdata = {
                game : this.selected_game,
                funcname : this.form.funcname,
                funcswitch: this.form.funcswitch
            }
            console.log(formdata)
            this.$store.dispatch('svrmgr/SetSvrSwitch', { user, formdata}).then((res) => {
                if (res.errorCode != 0) {
                    this.$message.error(`执行失败 :` + res.errorCode);
                } else {
                    this.$set(this.tableData, this.idx, this.form);
                    this.$message.success(`执行成功`);
                }
                //this.editVisible = false;
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });

            this.editVisible = false;
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
