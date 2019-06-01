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
                <!--
                <span class="demonstration">游戏ID</span>
                <xdh-select v-model="selected_game"> </xdh-select>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                <el-col class="line" :span="20"></el-col>
                <el-button type="primary"  @click="handleAdd">新增<i class="el-icon-upload el-icon--right"></i></el-button>-->

            </div>
            <el-table :data="data" border class="table">
                <el-table-column prop="game" label="游戏" width="100">
                </el-table-column>
                <el-table-column prop="jsonvalue" label="配置" width="600">
                    <template slot-scope="scope">
                        <json-viewer :value="scope.row.jsonvalue" :expand-depth=5 boxed ></json-viewer>
                    </template>
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
                <el-form-item label="配置">
                    <el-input type="textarea" :rows="10" placeholder="请输入内容" v-model="form.jsonvalue"></el-input>
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
import {compileStr, uncompileStr} from '@/utils/index'
export default {
    components: { XdhSelect },
    name: 'settings',
    data() {
        return {
            tableData: [],
            cur_page: 1,
            is_search: false,
            selected_game : '',
            editVisible: false,
            form: {
                game:'',
                jsonvalue: ''           
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
            this.handleSearch();
        },

        getGameID() {
            return "11111"
            // let child = this.$refs.gameID
            // console.log(child.getValue())
            // return child.getValue()
        },
        handleSearch() {
            this.is_search = true;
            const user = this.$store.getters.name
            const page = this.$route.path
            const game = this.selected_game
            console.log(user, page, this.getGameID())
            console.log(this.selected_game)
            this.$store.dispatch('svrmgr/GetSvrCfg', { user, page, game}).then((res) => {
                console.log(res)
                if (res.errorCode != 0) {
                    this.$message.error('服务器返回失败:' + res.errorCode);
                } else {
                    const svalue = uncompileStr(res.data)
                    this.tableData = [{game:game, jsonvalue:svalue}]
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
                jsonvalue: ''
            }
            this.editVisible = true
        },
        handleEdit(index, row) {
            this.idx = index;
            const item = this.tableData[index];
            this.form = {
                game: this.selected_game,
                jsonvalue: item.jsonvalue
            }
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
            const user = this.$store.getters.name
            const formdata = {
                game : this.selected_game,
                jsonvalue : compileStr(this.form.jsonvalue)
            }
            // console.log("SetSvrCfg >>>>>>> ")
            // console.log(formdata)
            // // console.log(this.select_daterange)
            this.$store.dispatch('svrmgr/SetSvrCfg', { user, formdata}).then((res) => {
                this.$set(this.tableData, this.idx, this.form);
                //this.editVisible = false;
                this.$message.success(`执行成功`);
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
        width: 90%;
        font-size: 14px;
    }
    .red{
        color: #ff0000;
    }
    .mr10{
        margin-right: 10px;
    }
</style>
