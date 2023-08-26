# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: index/genesis.proto, index/index.proto, index/params.proto, index/query.proto, index/sql_engine.proto, index/tx.proto
# plugin: python-betterproto
# This file has been @generated

from dataclasses import dataclass
from typing import (
    TYPE_CHECKING,
    Dict,
    List,
    Optional,
)

import betterproto
import grpclib
from betterproto.grpc.grpclib_server import ServiceBase

from ....cosmos.base.query import v1beta1 as ___cosmos_base_query_v1_beta1__


if TYPE_CHECKING:
    import grpclib.server
    from betterproto.grpc.grpclib_client import MetadataLike
    from grpclib.metadata import Deadline


class ColumnValueType(betterproto.Enum):
    InvalidColumn = 0
    IntColumn = 1
    UintColumn = 2
    FloatColumn = 3
    BoolColumn = 4
    StringColumn = 5
    BytesColumn = 6


class ArgumentVarType(betterproto.Enum):
    INT = 0
    UINT = 1
    FLOAT = 2
    BOOL = 3
    STRING = 4
    BYTES = 5


@dataclass(eq=False, repr=False)
class Argument(betterproto.Message):
    type: "ArgumentVarType" = betterproto.enum_field(1)
    """variable type"""

    value: str = betterproto.string_field(3)
    """variable value"""


@dataclass(eq=False, repr=False)
class RowData(betterproto.Message):
    columns: List[str] = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class ResultSet(betterproto.Message):
    id: str = betterproto.string_field(1)
    column_defs: List["ColumnDef"] = betterproto.message_field(2)
    rows: List["RowData"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class ColumnDef(betterproto.Message):
    column_name: str = betterproto.string_field(1)
    column_type: str = betterproto.string_field(2)
    comment: str = betterproto.string_field(3)
    column_value_type: "ColumnValueType" = betterproto.enum_field(4)


@dataclass(eq=False, repr=False)
class IndexDef(betterproto.Message):
    index_name: str = betterproto.string_field(1)
    index_type: str = betterproto.string_field(2)
    columns: List["ColumnDef"] = betterproto.message_field(4)
    is_primary_key: bool = betterproto.bool_field(5)
    parser: str = betterproto.string_field(6)
    comment: str = betterproto.string_field(7)


@dataclass(eq=False, repr=False)
class TableInfo(betterproto.Message):
    table_name: str = betterproto.string_field(1)
    db_name: str = betterproto.string_field(2)
    engine: str = betterproto.string_field(3)
    columns: List["ColumnDef"] = betterproto.message_field(4)
    indexes: List["IndexDef"] = betterproto.message_field(5)
    create_statement: str = betterproto.string_field(6)
    comment: str = betterproto.string_field(7)
    creator: str = betterproto.string_field(8)
    members: List["TableMember"] = betterproto.message_field(9)


@dataclass(eq=False, repr=False)
class TableMember(betterproto.Message):
    uid: str = betterproto.string_field(1)
    role: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class DatabaseInfo(betterproto.Message):
    db_name: str = betterproto.string_field(1)
    creator: str = betterproto.string_field(2)
    members: List["TableMember"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class MsgSchema(betterproto.Message):
    owner_address: str = betterproto.string_field(1)
    schema_name: str = betterproto.string_field(2)
    body: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class MsgSchemaResponse(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class MsgDoc(betterproto.Message):
    owner_address: str = betterproto.string_field(1)
    schema_name: str = betterproto.string_field(2)
    body: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class MsgDocResponse(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class SqlExecRequest(betterproto.Message):
    uid: str = betterproto.string_field(1)
    sql: str = betterproto.string_field(2)
    arguments: List["Argument"] = betterproto.message_field(3)


@dataclass(eq=False, repr=False)
class SqlExecResponse(betterproto.Message):
    results: List["ResultSet"] = betterproto.message_field(1)
    took_times: float = betterproto.float_field(2)


@dataclass(eq=False, repr=False)
class SqlGrantRequest(betterproto.Message):
    uid: str = betterproto.string_field(1)
    on_table: str = betterproto.string_field(2)
    to_uid: str = betterproto.string_field(3)
    role: str = betterproto.string_field(4)
    on_database: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class SqlGrantResponse(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class Params(betterproto.Message):
    """Params defines the parameters for the module."""

    pass


@dataclass(eq=False, repr=False)
class Schema(betterproto.Message):
    owner_address: str = betterproto.string_field(1)
    schema_name: str = betterproto.string_field(2)
    body: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class Doc(betterproto.Message):
    owner_address: str = betterproto.string_field(1)
    schema_name: str = betterproto.string_field(2)
    body: str = betterproto.string_field(3)


@dataclass(eq=False, repr=False)
class ListSchemaRequest(betterproto.Message):
    """ListSchemaRequest is request type for the Query RPC method."""

    pagination: "___cosmos_base_query_v1_beta1__.PageRequest" = (
        betterproto.message_field(1)
    )


@dataclass(eq=False, repr=False)
class ListSchemaResponse(betterproto.Message):
    """ListSchemaResponse is response type for the Query RPC method."""

    schemas: List["Schema"] = betterproto.message_field(1)
    pagination: "___cosmos_base_query_v1_beta1__.PageResponse" = (
        betterproto.message_field(2)
    )


@dataclass(eq=False, repr=False)
class GetSchemaRequest(betterproto.Message):
    """GetSchemaRequest is request type for the Query RPC method."""

    schema_name: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class GetSchemaResponse(betterproto.Message):
    """GetSchemaResponse is response type for the Query RPC method."""

    schema: "Schema" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class GetDocRequest(betterproto.Message):
    """GetDocsRequest is request type for the Query RPC method."""

    schema_name: str = betterproto.string_field(1)
    doc_id: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class GetDocResponse(betterproto.Message):
    """GetDocsResponse is response type for the Query RPC method."""

    doc: "Doc" = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class SearchDocRequest(betterproto.Message):
    """GetDocsRequest is request type for the Query RPC method."""

    schema_name: str = betterproto.string_field(1)
    query: str = betterproto.string_field(2)
    query_field: List[str] = betterproto.string_field(3)
    aggs_field: List[str] = betterproto.string_field(4)
    filter: "SearchFilters" = betterproto.message_field(5)
    order_by: str = betterproto.string_field(6)
    limit: int = betterproto.int64_field(7)
    page: int = betterproto.int64_field(8)
    factor: "FieldValueFactor" = betterproto.message_field(9)


@dataclass(eq=False, repr=False)
class SearchDocResponse(betterproto.Message):
    """GetDocsResponse is response type for the Query RPC method."""

    search_time: int = betterproto.int64_field(1)
    schema_name: str = betterproto.string_field(2)
    meta: "Meta" = betterproto.message_field(3)
    items: List["DocItem"] = betterproto.message_field(4)
    facet: Dict[str, "SearchFilters"] = betterproto.map_field(
        5, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class Meta(betterproto.Message):
    current_page: int = betterproto.int64_field(1)
    total_pages: int = betterproto.int64_field(2)
    total_results: int = betterproto.int64_field(3)
    size: int = betterproto.int64_field(4)
    sorted_by: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class DocItem(betterproto.Message):
    highlight: Dict[str, "Strings"] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )
    data: str = betterproto.string_field(2)
    score: float = betterproto.double_field(3)


@dataclass(eq=False, repr=False)
class Strings(betterproto.Message):
    strings: List[str] = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class SearchFilters(betterproto.Message):
    search_filter: List["SearchFilter"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class SearchFilter(betterproto.Message):
    type: str = betterproto.string_field(1)
    field: str = betterproto.string_field(2)
    value: str = betterproto.string_field(3)
    from_: float = betterproto.float_field(4)
    to: float = betterproto.float_field(5)
    doc_count: int = betterproto.int64_field(6)


@dataclass(eq=False, repr=False)
class FieldValueFactor(betterproto.Message):
    field: str = betterproto.string_field(1)
    factor: float = betterproto.double_field(2)
    missing: float = betterproto.double_field(3)
    weight: float = betterproto.double_field(4)
    modifier: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class QueryParamsRequest(betterproto.Message):
    """QueryParamsRequest is request type for the Query/Params RPC method."""

    pass


@dataclass(eq=False, repr=False)
class QueryParamsResponse(betterproto.Message):
    """
    QueryParamsResponse is response type for the Query/Params RPC method.
    """

    params: "Params" = betterproto.message_field(1)
    """params holds all the parameters of this module."""


@dataclass(eq=False, repr=False)
class SqlQueryRequest(betterproto.Message):
    sql: str = betterproto.string_field(1)
    arguments: List["Argument"] = betterproto.message_field(2)


@dataclass(eq=False, repr=False)
class SqlQueryResponse(betterproto.Message):
    results: List["ResultSet"] = betterproto.message_field(1)
    took_times: float = betterproto.float_field(2)


@dataclass(eq=False, repr=False)
class SimpleSqlQueryResponse(betterproto.Message):
    result: List["SimpleSqlQueryResponseResultSet"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class SimpleSqlQueryResponseResultSet(betterproto.Message):
    row: Dict[str, "SimpleSqlQueryResponseRowValue"] = betterproto.map_field(
        1, betterproto.TYPE_STRING, betterproto.TYPE_MESSAGE
    )


@dataclass(eq=False, repr=False)
class SimpleSqlQueryResponseRowValue(betterproto.Message):
    value: str = betterproto.string_field(1)
    column_value_type: "ColumnValueType" = betterproto.enum_field(2)


@dataclass(eq=False, repr=False)
class SqlListTablesRequest(betterproto.Message):
    keyword: str = betterproto.string_field(1)
    uid: str = betterproto.string_field(2)
    page: int = betterproto.int32_field(3)
    page_size: int = betterproto.int32_field(4)
    database: str = betterproto.string_field(5)


@dataclass(eq=False, repr=False)
class SqlListTablesResponse(betterproto.Message):
    tables: List["TableInfo"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class SqlListDatabasesRequest(betterproto.Message):
    pass


@dataclass(eq=False, repr=False)
class SqlListDatabasesResponse(betterproto.Message):
    databases: List["DatabaseInfo"] = betterproto.message_field(1)


@dataclass(eq=False, repr=False)
class ShowCreateTableRequest(betterproto.Message):
    database_name: str = betterproto.string_field(1)
    table_name: str = betterproto.string_field(2)


@dataclass(eq=False, repr=False)
class ShowCreateTableResponse(betterproto.Message):
    schema: str = betterproto.string_field(1)


@dataclass(eq=False, repr=False)
class GenesisState(betterproto.Message):
    """GenesisState defines the index module's genesis state."""

    schema_list: List["Schema"] = betterproto.message_field(1)


class MsgStub(betterproto.ServiceStub):
    async def manage_schema(
        self,
        msg_schema: "MsgSchema",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "MsgSchemaResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Msg/ManageSchema",
            msg_schema,
            MsgSchemaResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def manage_doc(
        self,
        msg_doc: "MsgDoc",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "MsgDocResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Msg/ManageDoc",
            msg_doc,
            MsgDocResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def sql_exec(
        self,
        sql_exec_request: "SqlExecRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SqlExecResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Msg/SQLExec",
            sql_exec_request,
            SqlExecResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def sql_grant(
        self,
        sql_grant_request: "SqlGrantRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SqlGrantResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Msg/SQLGrant",
            sql_grant_request,
            SqlGrantResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )


class QueryStub(betterproto.ServiceStub):
    async def list_schema(
        self,
        list_schema_request: "ListSchemaRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "ListSchemaResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/ListSchema",
            list_schema_request,
            ListSchemaResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def get_schema(
        self,
        get_schema_request: "GetSchemaRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "GetSchemaResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/GetSchema",
            get_schema_request,
            GetSchemaResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def get_doc(
        self,
        get_doc_request: "GetDocRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "GetDocResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/GetDoc",
            get_doc_request,
            GetDocResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def search_doc(
        self,
        search_doc_request: "SearchDocRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SearchDocResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/SearchDoc",
            search_doc_request,
            SearchDocResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def params(
        self,
        query_params_request: "QueryParamsRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "QueryParamsResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/Params",
            query_params_request,
            QueryParamsResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def sql_query(
        self,
        sql_query_request: "SqlQueryRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SqlQueryResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/SQLQuery",
            sql_query_request,
            SqlQueryResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def simple_sql_query(
        self,
        sql_query_request: "SqlQueryRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SimpleSqlQueryResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/SimpleSQLQuery",
            sql_query_request,
            SimpleSqlQueryResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def sql_list_tables(
        self,
        sql_list_tables_request: "SqlListTablesRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SqlListTablesResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/SQLListTables",
            sql_list_tables_request,
            SqlListTablesResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def sql_list_databases(
        self,
        sql_list_databases_request: "SqlListDatabasesRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "SqlListDatabasesResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/SQLListDatabases",
            sql_list_databases_request,
            SqlListDatabasesResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )

    async def show_create_table(
        self,
        show_create_table_request: "ShowCreateTableRequest",
        *,
        timeout: Optional[float] = None,
        deadline: Optional["Deadline"] = None,
        metadata: Optional["MetadataLike"] = None
    ) -> "ShowCreateTableResponse":
        return await self._unary_unary(
            "/blockved.glitterchain.index.Query/ShowCreateTable",
            show_create_table_request,
            ShowCreateTableResponse,
            timeout=timeout,
            deadline=deadline,
            metadata=metadata,
        )


class MsgBase(ServiceBase):
    async def manage_schema(self, msg_schema: "MsgSchema") -> "MsgSchemaResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def manage_doc(self, msg_doc: "MsgDoc") -> "MsgDocResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def sql_exec(self, sql_exec_request: "SqlExecRequest") -> "SqlExecResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def sql_grant(
        self, sql_grant_request: "SqlGrantRequest"
    ) -> "SqlGrantResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_manage_schema(
        self, stream: "grpclib.server.Stream[MsgSchema, MsgSchemaResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.manage_schema(request)
        await stream.send_message(response)

    async def __rpc_manage_doc(
        self, stream: "grpclib.server.Stream[MsgDoc, MsgDocResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.manage_doc(request)
        await stream.send_message(response)

    async def __rpc_sql_exec(
        self, stream: "grpclib.server.Stream[SqlExecRequest, SqlExecResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.sql_exec(request)
        await stream.send_message(response)

    async def __rpc_sql_grant(
        self, stream: "grpclib.server.Stream[SqlGrantRequest, SqlGrantResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.sql_grant(request)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/blockved.glitterchain.index.Msg/ManageSchema": grpclib.const.Handler(
                self.__rpc_manage_schema,
                grpclib.const.Cardinality.UNARY_UNARY,
                MsgSchema,
                MsgSchemaResponse,
            ),
            "/blockved.glitterchain.index.Msg/ManageDoc": grpclib.const.Handler(
                self.__rpc_manage_doc,
                grpclib.const.Cardinality.UNARY_UNARY,
                MsgDoc,
                MsgDocResponse,
            ),
            "/blockved.glitterchain.index.Msg/SQLExec": grpclib.const.Handler(
                self.__rpc_sql_exec,
                grpclib.const.Cardinality.UNARY_UNARY,
                SqlExecRequest,
                SqlExecResponse,
            ),
            "/blockved.glitterchain.index.Msg/SQLGrant": grpclib.const.Handler(
                self.__rpc_sql_grant,
                grpclib.const.Cardinality.UNARY_UNARY,
                SqlGrantRequest,
                SqlGrantResponse,
            ),
        }


class QueryBase(ServiceBase):
    async def list_schema(
        self, list_schema_request: "ListSchemaRequest"
    ) -> "ListSchemaResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def get_schema(
        self, get_schema_request: "GetSchemaRequest"
    ) -> "GetSchemaResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def get_doc(self, get_doc_request: "GetDocRequest") -> "GetDocResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def search_doc(
        self, search_doc_request: "SearchDocRequest"
    ) -> "SearchDocResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def params(
        self, query_params_request: "QueryParamsRequest"
    ) -> "QueryParamsResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def sql_query(
        self, sql_query_request: "SqlQueryRequest"
    ) -> "SqlQueryResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def simple_sql_query(
        self, sql_query_request: "SqlQueryRequest"
    ) -> "SimpleSqlQueryResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def sql_list_tables(
        self, sql_list_tables_request: "SqlListTablesRequest"
    ) -> "SqlListTablesResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def sql_list_databases(
        self, sql_list_databases_request: "SqlListDatabasesRequest"
    ) -> "SqlListDatabasesResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def show_create_table(
        self, show_create_table_request: "ShowCreateTableRequest"
    ) -> "ShowCreateTableResponse":
        raise grpclib.GRPCError(grpclib.const.Status.UNIMPLEMENTED)

    async def __rpc_list_schema(
        self, stream: "grpclib.server.Stream[ListSchemaRequest, ListSchemaResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.list_schema(request)
        await stream.send_message(response)

    async def __rpc_get_schema(
        self, stream: "grpclib.server.Stream[GetSchemaRequest, GetSchemaResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.get_schema(request)
        await stream.send_message(response)

    async def __rpc_get_doc(
        self, stream: "grpclib.server.Stream[GetDocRequest, GetDocResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.get_doc(request)
        await stream.send_message(response)

    async def __rpc_search_doc(
        self, stream: "grpclib.server.Stream[SearchDocRequest, SearchDocResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.search_doc(request)
        await stream.send_message(response)

    async def __rpc_params(
        self, stream: "grpclib.server.Stream[QueryParamsRequest, QueryParamsResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.params(request)
        await stream.send_message(response)

    async def __rpc_sql_query(
        self, stream: "grpclib.server.Stream[SqlQueryRequest, SqlQueryResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.sql_query(request)
        await stream.send_message(response)

    async def __rpc_simple_sql_query(
        self, stream: "grpclib.server.Stream[SqlQueryRequest, SimpleSqlQueryResponse]"
    ) -> None:
        request = await stream.recv_message()
        response = await self.simple_sql_query(request)
        await stream.send_message(response)

    async def __rpc_sql_list_tables(
        self,
        stream: "grpclib.server.Stream[SqlListTablesRequest, SqlListTablesResponse]",
    ) -> None:
        request = await stream.recv_message()
        response = await self.sql_list_tables(request)
        await stream.send_message(response)

    async def __rpc_sql_list_databases(
        self,
        stream: "grpclib.server.Stream[SqlListDatabasesRequest, SqlListDatabasesResponse]",
    ) -> None:
        request = await stream.recv_message()
        response = await self.sql_list_databases(request)
        await stream.send_message(response)

    async def __rpc_show_create_table(
        self,
        stream: "grpclib.server.Stream[ShowCreateTableRequest, ShowCreateTableResponse]",
    ) -> None:
        request = await stream.recv_message()
        response = await self.show_create_table(request)
        await stream.send_message(response)

    def __mapping__(self) -> Dict[str, grpclib.const.Handler]:
        return {
            "/blockved.glitterchain.index.Query/ListSchema": grpclib.const.Handler(
                self.__rpc_list_schema,
                grpclib.const.Cardinality.UNARY_UNARY,
                ListSchemaRequest,
                ListSchemaResponse,
            ),
            "/blockved.glitterchain.index.Query/GetSchema": grpclib.const.Handler(
                self.__rpc_get_schema,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetSchemaRequest,
                GetSchemaResponse,
            ),
            "/blockved.glitterchain.index.Query/GetDoc": grpclib.const.Handler(
                self.__rpc_get_doc,
                grpclib.const.Cardinality.UNARY_UNARY,
                GetDocRequest,
                GetDocResponse,
            ),
            "/blockved.glitterchain.index.Query/SearchDoc": grpclib.const.Handler(
                self.__rpc_search_doc,
                grpclib.const.Cardinality.UNARY_UNARY,
                SearchDocRequest,
                SearchDocResponse,
            ),
            "/blockved.glitterchain.index.Query/Params": grpclib.const.Handler(
                self.__rpc_params,
                grpclib.const.Cardinality.UNARY_UNARY,
                QueryParamsRequest,
                QueryParamsResponse,
            ),
            "/blockved.glitterchain.index.Query/SQLQuery": grpclib.const.Handler(
                self.__rpc_sql_query,
                grpclib.const.Cardinality.UNARY_UNARY,
                SqlQueryRequest,
                SqlQueryResponse,
            ),
            "/blockved.glitterchain.index.Query/SimpleSQLQuery": grpclib.const.Handler(
                self.__rpc_simple_sql_query,
                grpclib.const.Cardinality.UNARY_UNARY,
                SqlQueryRequest,
                SimpleSqlQueryResponse,
            ),
            "/blockved.glitterchain.index.Query/SQLListTables": grpclib.const.Handler(
                self.__rpc_sql_list_tables,
                grpclib.const.Cardinality.UNARY_UNARY,
                SqlListTablesRequest,
                SqlListTablesResponse,
            ),
            "/blockved.glitterchain.index.Query/SQLListDatabases": grpclib.const.Handler(
                self.__rpc_sql_list_databases,
                grpclib.const.Cardinality.UNARY_UNARY,
                SqlListDatabasesRequest,
                SqlListDatabasesResponse,
            ),
            "/blockved.glitterchain.index.Query/ShowCreateTable": grpclib.const.Handler(
                self.__rpc_show_create_table,
                grpclib.const.Cardinality.UNARY_UNARY,
                ShowCreateTableRequest,
                ShowCreateTableResponse,
            ),
        }