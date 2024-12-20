// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

export interface AnyRequest {}

export interface Category {
  id?: number
  createdAt?: string
  updatedAt?: string
  cid?: string
  depth?: number
  title: string
  order?: number
  path?: string
}

export interface Announce {
  iD: number
  createdAt: string
  updatedAt: string
  title: string
  content: string
  date: string
}

export interface AnnouncesData {
  announces: Array<Announce>
}

export interface Collection {
  cid: string
  title: string
  link: string
  path: string
  categoryID?: string
  category: Category
  description: string

  favicon: File | null
  tags?: Array<string>
  // view?: number
  like?: boolean
}

export interface CollectionsData {
  category: Category
  groups: Array<Group>
}

export interface CollectionsData {
  category: Category
  groups: Array<Group>
}

export interface CollectionsDataResponse {
  datas: Array<CollectionsData>
}
export interface CollectionResponse {
  data: Collection
}

export interface Comment {
  iD: number
  createdAt: string
  updatedAt: string
  nickname: string
  content: string
  date: string
  replies: Array<Comment>
  parent_id?: number
}

export interface CommentRequest {
  nickname: string
  content: string
}

export interface CommentsResponse {
  comments: Array<Comment>
}

export interface CollectionFilter {
  categories?: Array<string>
  page: number
  page_size?: number
  keyword: string
}

export interface CollectionUpdateParams {
  cid: string
  title: string
  link: string
  description?: string
  category_path?: string
  tags?: string
  favicon?: File | null
}

export interface CollectionsResponse {
  count: number
  totalPages: number
  collections: Array<Collection>
}
export interface Group {
  category: Category
  collections: Array<Collection>
  show: boolean
}

export interface IDRequest {
  id: string
}

export interface IDResponse {
  id: string
}

export interface Nav {
  cid?: string
  title: string
  full_title: string
  order?: number
  path?: string
  children?: Nav[]
}

export interface NavsResponse {
  navs: Array<Nav>
}

export interface TagRequest {
  tags: Array<string>
}

export interface TagsResponse {
  tags: Array<string>
}

export interface SiteStats {
  last_day_visitors: number // 昨日访问者数量
  total_visitors: number // 总访问者数量
  links_count: number // 链接数量
}

export interface Statistics {
  iD: number
  createdAt: string
  updatedAt: string
  ip: string // 访问者的 IP 地址
  date: string // 访问日期和时间
}

export interface User {
  iD: number
  createdAt: string
  updatedAt: string
  username: string
  email?: string
  password: string
  nickname?: string
}

export interface UserResponse {}
export interface UserResponseParams {}

export interface AuthResponse {
  username: string
  email: string
  nickname: string
  role: string
  token: string
}

export interface AuthResponseParams {}
