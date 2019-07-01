<template>
    <div class="table">
        <div class="app-container">
            <div class="handle-box">
                <el-row :gutter="10">
                    <el-col :span="4">
                        <xdh-select v-model="selected_game"> </xdh-select>
                    </el-col>
                </el-row>
            </div>
            <div class="handle-box">
                <el-row :gutter="10">
                    <el-col :span="4">
                        <el-input placeholder="请输入关键字" v-model="input_key"> </el-input>
                    </el-col>
                    <el-col :span="4">
                        <el-select v-model="input_ktype" placeholder="请选择" >
                            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                                <span style="float: left">{{ item.label }}</span>
                                <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                            </el-option>
                        </el-select>
                    </el-col>
                    <el-col :span="2" >
                        <el-button type="primary"  @click="handleSearch">搜索<i class="el-icon-upload el-icon--right"></i></el-button>
                    </el-col>
                </el-row>
            </div>

            <el-table :data="data" border class="table">
                <el-table-column prop="jsonvalue" label="配置" width="800">
                    <template slot-scope="scope">
                        <json-viewer :value="scope.row.jsonvalue" :expand-depth=5 boxed copyable ></json-viewer>
                    </template>
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="1">
                </el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="50%">
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="配置">
                    <template>
                    <div>
                        <vue-json-editor v-model="form.jsonvalue"></vue-json-editor>
                    <div>
                    </template>
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
import vueJsonEditor from 'vue-json-editor'
import XdhSelect from '../svrmgr/xdhselect'
import {isJSON} from '@/utils/index'
export default {
    components: { vueJsonEditor, XdhSelect },
    name: 'user',
    data() {
        return {
            options:[
                {
                    value: 'id',
                    label: 'id'
                },
                {
                    value: 'nick',
                    label: 'nick'
                },
                {
                    value: 'openid',
                    label: 'openid'
                }
            ],
            tableData: [],
            cur_page: 1,
            is_search: false,
            editVisible: false,
            selected_game : '',
            actioning : 'get',
            input_key : '',
            input_ktype : 'id',
            form: {
                game: this.selected_game,
                key : '',
                jsonvalue: null
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
            const key = this.input_key
            const ktype = this.input_ktype
            console.log(this.selected_game)
            this.tableData = []
            this.$store.dispatch('svrmgr/GetPlayerInfo', { user, page, game, key, ktype}).then((res) => {
                console.log(res)
                if (res.errorCode != 0) {
                    this.$message.error('服务器返回失败:' + res.errorCode)
                } else {
                    if (res.data) {
                        this.tableData = [{jsonvalue:res.data}]
                    } else {
                        this.$message.error('没有数据')
                    }
                } 
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });
        },
        handleEdit(index, row) {
            this.idx = index
            if (this.selected_game === '') {
                this.$message.success(`请先选择游戏`)
                return
            }
            // var sjsonvalue = JSON.stringify(this.tableData[index].jsonvalue)
            var jsonobj = this.tableData[index].jsonvalue
            // console.log("--------------------------------")
            console.log(jsonobj["_id"])
            // console.log(sjsonvalue)
            // console.log("--------------------------------")
            this.form = {
                game: this.selected_game,
                key : jsonobj["_id"],
                jsonvalue: jsonobj
            }
            console.log(this.form.jsonvalue)

            this.actioning = 'set'
            this.editVisible = true;
        },
        // 保存编辑
        saveEdit() {
            var sjsonvalue = JSON.stringify(this.form.jsonvalue)
            if (!isJSON(sjsonvalue)) {
                this.$message.error(`Json 格式错误`);
                return
            }

            const user = this.$store.getters.name
            const formdata = {
                game : this.selected_game,
                uid : this.form.key,
                data : sjsonvalue
            }
            console.log(formdata)
            this.$store.dispatch('svrmgr/EditPlayerInfo', { user, formdata}).then((res) => {
                if (res.errorCode != 0) {
                    this.$message.error(`执行失败 :` + res.errorCode);
                } else {
                    this.$message.success(`执行成功`);
                }
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
