import { Component, Inject, OnInit } from '@angular/core'
import { EoNgFeedbackModalService } from 'eo-ng-feedback'
import { NzModalRef } from 'ng-zorro-antd/modal'
import { MODAL_NORMAL_SIZE, MODAL_SMALL_SIZE } from '../../../constant/app.config'
import { EmptyHttpResponse } from '../../../constant/type'
import { API_URL, ApiService } from '../../../service/api.service'
import { BaseInfoService } from '../../../service/base-info.service'
import { EoNgMessageService } from '../../../service/eo-ng-message.service'
import { EoNgNavigationService } from '../../../service/eo-ng-navigation.service'
import { PluginConfigComponent } from '../config/config.component'
import { PluginMessage } from '../types/types'
import { MarkdownService } from 'ngx-markdown'

@Component({
  selector: 'eo-ng-plugin-message',
  template: `
  <header class="my-btnybase ml-btnbase mr-btnrbase">
    <div class="flex justify-between  mb-btnybase items-center">
      <div class="flex">
      <img class="mr-btnrbase" [src]="icon" alt="plugin icon" width="64px" height="50px">

        <div>
          <p class="text-[18px] font-medium">{{title}}</p>
          <p>{{enable? '启用':'未启用'}}</p>
        </div>
      </div>
      <div>
      <button
        *ngIf="!enable"
        class="ml-btnybase ant-btn-primary"
        eo-ng-button
        (click)="enablePlugin()"
      >
        启用
      </button>
      <button
        *ngIf="enable"
        class="ml-btnybase"
        eo-ng-button
        (click)="disablePluginModal(false)"
      >
        停用
      </button>
      <button
      *ngIf="!enable && uninstall"
        class="ml-btnybase"
        eo-ng-button
        (click)="disablePluginModal(true)"
      >
        卸载
      </button>
      </div>
    </div>
    <p>{{resume}}</p>

  </header>
  <section class="block ml-btnbase mr-btnrbase p-btnbase markdown-block">
    <markdown [src]="getMd()" [srcRelativeLink]="true"   (error)="onError($event)"></markdown>
  </section>
  `,
  styles: [
    `
    .markdown-block{
      border:1px solid var(--border-color);
    }
    :host ::ng-deep{
      img{
        max-width:100%
      }
    }`
  ]
})
export class PluginMessageComponent implements OnInit {
  title:string = ''
  resume:string = ''
  icon:string = '' || './assets/default-plugin-icon.svg'
  enable:boolean = false
  uninstall:boolean = false
  pluginId:string = ''
  modalRef:NzModalRef|undefined
  mdFileName:string = ''
  constructor (private message:EoNgMessageService,
    private modalService:EoNgFeedbackModalService,
    private api:ApiService, private baseInfo:BaseInfoService,
    private appConfigService: EoNgNavigationService,
    private markdownService: MarkdownService,
    @Inject(API_URL) public urlPrefix:string) {
    this.appConfigService.reqFlashBreadcrumb([
      { title: '企业插件', routerLink: ['/', 'plugin', 'list', ''] },
      { title: '插件详情' }
    ])
  }

  ngOnInit (): void {
    this.pluginId = this.baseInfo.allParamsInfo.pluginId
    this.mdFileName = this.baseInfo.allParamsInfo.mdFileName
    this.getPluginDetail()
    this.markdownService.renderer.link = (href, title, text) => {
      let html = ''
      if (href && /^(?![http])[.]*/.test(href!) && /^(?![#])[.]*/.test(href!) && href.includes('.md')) {
        html = `<a href="plugin/message/${this.pluginId}/${href}">${text}</a>`
      } else if (href && /^(?![http])[.]*/.test(href!) && /^(?![#])[.]*/.test(href!)) {
        html = `<a href="plugin/message/${this.pluginId}/${href}">${text}</a>`
      } else {
        html = '<a  role="link"  tabindex="0" target="_blank" rel="nofollow noopener noreferrer" href="' + href + '">' + text + '</a>'
      }

      return html
    }

    this.markdownService.renderer.image = (src, title, alt) => {
      let html
      if (src && /^(?![http])[.]*/.test(src!)) {
        html = `<image src="${this.urlPrefix}plugin/info/${this.pluginId}/resource/${src}" alt=${alt}/>`
      } else {
        html = `<image src="${src}" alt=${alt}/>`
      }
      return html
    }
  }

  getPluginDetail () {
    this.api.get('system/plugin/info', { id: this.pluginId })
      .subscribe((resp:{code:number, data:{plugin:PluginMessage}, msg:string}) => {
        if (resp.code === 0) {
          this.title = resp.data.plugin.cname
          this.resume = resp.data.plugin.resume
          this.icon = resp.data.plugin.icon
          this.enable = resp.data.plugin.enable
          this.uninstall = resp.data.plugin.uninstall
        }
      })
  }

  getMd () {
    return `../../plugin/info/${this.pluginId}/${this.mdFileName}'`
  }

  onError (value:any) {
    console.error('解析md文档出现问题', value)
  }

  enablePlugin () {
    this.modalService.create({
      nzTitle: '启用',
      nzWidth: MODAL_NORMAL_SIZE,
      nzContent: PluginConfigComponent,
      nzComponentParams: {

      },
      nzOkText: '确定',
      nzCancelText: '取消',
      nzOnOk: (component:PluginConfigComponent) => {
        component.enablePlugin()
        return false
      }
    })
  }

  disablePluginModal (deletePlugin:boolean) {
    this.modalRef = this.modalService.create({
      nzTitle: deletePlugin ? '卸载' : '停用',
      nzContent: `该插件${deletePlugin ? '卸载后将无法找回' : '停用后将无法再使用'}，请确认是否要${deletePlugin ? '卸载' : '停用'}？`,
      nzClosable: true,
      nzClassName: 'delete-modal',
      nzWidth: MODAL_SMALL_SIZE,
      nzOkText: '确定',
      nzOkDanger: true,
      nzCancelText: '取消',
      nzOnOk: () => {
        if (deletePlugin) {
          this.deletePlugin()
        } else {
          this.disablePlugin()
        }
        return false
      }
    })
  }

  // 禁用插件
  disablePlugin () {
    this.api.post('system/plugin/disable', { id: this.pluginId }).subscribe((resp:EmptyHttpResponse) => {
      if (resp.code === 0) {
        this.message.success(resp.msg || '禁用成功')
        this.modalRef?.close()
      }
    })
  }

  // 卸载插件
  deletePlugin () {
    this.api.post('system/plugin/uninstall', { id: this.pluginId }).subscribe((resp:EmptyHttpResponse) => {
      if (resp.code === 0) {
        this.message.success(resp.msg || '卸载成功')
        this.modalRef?.close()
      }
    })
  }
}
