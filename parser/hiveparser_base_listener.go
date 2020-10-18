// Code generated from HiveParser.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // HiveParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHiveParserListener is a complete listener for a parse tree produced by HiveParser.
type BaseHiveParserListener struct{}

var _ HiveParserListener = &BaseHiveParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHiveParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHiveParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHiveParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHiveParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseHiveParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseHiveParserListener) ExitStatement(ctx *StatementContext) {}

// EnterExplainStatement is called when production explainStatement is entered.
func (s *BaseHiveParserListener) EnterExplainStatement(ctx *ExplainStatementContext) {}

// ExitExplainStatement is called when production explainStatement is exited.
func (s *BaseHiveParserListener) ExitExplainStatement(ctx *ExplainStatementContext) {}

// EnterExplainOption is called when production explainOption is entered.
func (s *BaseHiveParserListener) EnterExplainOption(ctx *ExplainOptionContext) {}

// ExitExplainOption is called when production explainOption is exited.
func (s *BaseHiveParserListener) ExitExplainOption(ctx *ExplainOptionContext) {}

// EnterVectorizationOnly is called when production vectorizationOnly is entered.
func (s *BaseHiveParserListener) EnterVectorizationOnly(ctx *VectorizationOnlyContext) {}

// ExitVectorizationOnly is called when production vectorizationOnly is exited.
func (s *BaseHiveParserListener) ExitVectorizationOnly(ctx *VectorizationOnlyContext) {}

// EnterVectorizatonDetail is called when production vectorizatonDetail is entered.
func (s *BaseHiveParserListener) EnterVectorizatonDetail(ctx *VectorizatonDetailContext) {}

// ExitVectorizatonDetail is called when production vectorizatonDetail is exited.
func (s *BaseHiveParserListener) ExitVectorizatonDetail(ctx *VectorizatonDetailContext) {}

// EnterExecStatement is called when production execStatement is entered.
func (s *BaseHiveParserListener) EnterExecStatement(ctx *ExecStatementContext) {}

// ExitExecStatement is called when production execStatement is exited.
func (s *BaseHiveParserListener) ExitExecStatement(ctx *ExecStatementContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseHiveParserListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseHiveParserListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterReplicationClause is called when production replicationClause is entered.
func (s *BaseHiveParserListener) EnterReplicationClause(ctx *ReplicationClauseContext) {}

// ExitReplicationClause is called when production replicationClause is exited.
func (s *BaseHiveParserListener) ExitReplicationClause(ctx *ReplicationClauseContext) {}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseHiveParserListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseHiveParserListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseHiveParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseHiveParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterReplDumpStatement is called when production replDumpStatement is entered.
func (s *BaseHiveParserListener) EnterReplDumpStatement(ctx *ReplDumpStatementContext) {}

// ExitReplDumpStatement is called when production replDumpStatement is exited.
func (s *BaseHiveParserListener) ExitReplDumpStatement(ctx *ReplDumpStatementContext) {}

// EnterReplLoadStatement is called when production replLoadStatement is entered.
func (s *BaseHiveParserListener) EnterReplLoadStatement(ctx *ReplLoadStatementContext) {}

// ExitReplLoadStatement is called when production replLoadStatement is exited.
func (s *BaseHiveParserListener) ExitReplLoadStatement(ctx *ReplLoadStatementContext) {}

// EnterReplConfigs is called when production replConfigs is entered.
func (s *BaseHiveParserListener) EnterReplConfigs(ctx *ReplConfigsContext) {}

// ExitReplConfigs is called when production replConfigs is exited.
func (s *BaseHiveParserListener) ExitReplConfigs(ctx *ReplConfigsContext) {}

// EnterReplConfigsList is called when production replConfigsList is entered.
func (s *BaseHiveParserListener) EnterReplConfigsList(ctx *ReplConfigsListContext) {}

// ExitReplConfigsList is called when production replConfigsList is exited.
func (s *BaseHiveParserListener) ExitReplConfigsList(ctx *ReplConfigsListContext) {}

// EnterReplStatusStatement is called when production replStatusStatement is entered.
func (s *BaseHiveParserListener) EnterReplStatusStatement(ctx *ReplStatusStatementContext) {}

// ExitReplStatusStatement is called when production replStatusStatement is exited.
func (s *BaseHiveParserListener) ExitReplStatusStatement(ctx *ReplStatusStatementContext) {}

// EnterDdlStatement is called when production ddlStatement is entered.
func (s *BaseHiveParserListener) EnterDdlStatement(ctx *DdlStatementContext) {}

// ExitDdlStatement is called when production ddlStatement is exited.
func (s *BaseHiveParserListener) ExitDdlStatement(ctx *DdlStatementContext) {}

// EnterIfExists is called when production ifExists is entered.
func (s *BaseHiveParserListener) EnterIfExists(ctx *IfExistsContext) {}

// ExitIfExists is called when production ifExists is exited.
func (s *BaseHiveParserListener) ExitIfExists(ctx *IfExistsContext) {}

// EnterRestrictOrCascade is called when production restrictOrCascade is entered.
func (s *BaseHiveParserListener) EnterRestrictOrCascade(ctx *RestrictOrCascadeContext) {}

// ExitRestrictOrCascade is called when production restrictOrCascade is exited.
func (s *BaseHiveParserListener) ExitRestrictOrCascade(ctx *RestrictOrCascadeContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseHiveParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseHiveParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}

// EnterRewriteEnabled is called when production rewriteEnabled is entered.
func (s *BaseHiveParserListener) EnterRewriteEnabled(ctx *RewriteEnabledContext) {}

// ExitRewriteEnabled is called when production rewriteEnabled is exited.
func (s *BaseHiveParserListener) ExitRewriteEnabled(ctx *RewriteEnabledContext) {}

// EnterRewriteDisabled is called when production rewriteDisabled is entered.
func (s *BaseHiveParserListener) EnterRewriteDisabled(ctx *RewriteDisabledContext) {}

// ExitRewriteDisabled is called when production rewriteDisabled is exited.
func (s *BaseHiveParserListener) ExitRewriteDisabled(ctx *RewriteDisabledContext) {}

// EnterStoredAsDirs is called when production storedAsDirs is entered.
func (s *BaseHiveParserListener) EnterStoredAsDirs(ctx *StoredAsDirsContext) {}

// ExitStoredAsDirs is called when production storedAsDirs is exited.
func (s *BaseHiveParserListener) ExitStoredAsDirs(ctx *StoredAsDirsContext) {}

// EnterOrReplace is called when production orReplace is entered.
func (s *BaseHiveParserListener) EnterOrReplace(ctx *OrReplaceContext) {}

// ExitOrReplace is called when production orReplace is exited.
func (s *BaseHiveParserListener) ExitOrReplace(ctx *OrReplaceContext) {}

// EnterCreateDatabaseStatement is called when production createDatabaseStatement is entered.
func (s *BaseHiveParserListener) EnterCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) {}

// ExitCreateDatabaseStatement is called when production createDatabaseStatement is exited.
func (s *BaseHiveParserListener) ExitCreateDatabaseStatement(ctx *CreateDatabaseStatementContext) {}

// EnterDbLocation is called when production dbLocation is entered.
func (s *BaseHiveParserListener) EnterDbLocation(ctx *DbLocationContext) {}

// ExitDbLocation is called when production dbLocation is exited.
func (s *BaseHiveParserListener) ExitDbLocation(ctx *DbLocationContext) {}

// EnterDbProperties is called when production dbProperties is entered.
func (s *BaseHiveParserListener) EnterDbProperties(ctx *DbPropertiesContext) {}

// ExitDbProperties is called when production dbProperties is exited.
func (s *BaseHiveParserListener) ExitDbProperties(ctx *DbPropertiesContext) {}

// EnterDbPropertiesList is called when production dbPropertiesList is entered.
func (s *BaseHiveParserListener) EnterDbPropertiesList(ctx *DbPropertiesListContext) {}

// ExitDbPropertiesList is called when production dbPropertiesList is exited.
func (s *BaseHiveParserListener) ExitDbPropertiesList(ctx *DbPropertiesListContext) {}

// EnterSwitchDatabaseStatement is called when production switchDatabaseStatement is entered.
func (s *BaseHiveParserListener) EnterSwitchDatabaseStatement(ctx *SwitchDatabaseStatementContext) {}

// ExitSwitchDatabaseStatement is called when production switchDatabaseStatement is exited.
func (s *BaseHiveParserListener) ExitSwitchDatabaseStatement(ctx *SwitchDatabaseStatementContext) {}

// EnterDropDatabaseStatement is called when production dropDatabaseStatement is entered.
func (s *BaseHiveParserListener) EnterDropDatabaseStatement(ctx *DropDatabaseStatementContext) {}

// ExitDropDatabaseStatement is called when production dropDatabaseStatement is exited.
func (s *BaseHiveParserListener) ExitDropDatabaseStatement(ctx *DropDatabaseStatementContext) {}

// EnterDatabaseComment is called when production databaseComment is entered.
func (s *BaseHiveParserListener) EnterDatabaseComment(ctx *DatabaseCommentContext) {}

// ExitDatabaseComment is called when production databaseComment is exited.
func (s *BaseHiveParserListener) ExitDatabaseComment(ctx *DatabaseCommentContext) {}

// EnterCreateTableStatement is called when production createTableStatement is entered.
func (s *BaseHiveParserListener) EnterCreateTableStatement(ctx *CreateTableStatementContext) {}

// ExitCreateTableStatement is called when production createTableStatement is exited.
func (s *BaseHiveParserListener) ExitCreateTableStatement(ctx *CreateTableStatementContext) {}

// EnterTruncateTableStatement is called when production truncateTableStatement is entered.
func (s *BaseHiveParserListener) EnterTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// ExitTruncateTableStatement is called when production truncateTableStatement is exited.
func (s *BaseHiveParserListener) ExitTruncateTableStatement(ctx *TruncateTableStatementContext) {}

// EnterDropTableStatement is called when production dropTableStatement is entered.
func (s *BaseHiveParserListener) EnterDropTableStatement(ctx *DropTableStatementContext) {}

// ExitDropTableStatement is called when production dropTableStatement is exited.
func (s *BaseHiveParserListener) ExitDropTableStatement(ctx *DropTableStatementContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *BaseHiveParserListener) EnterAlterStatement(ctx *AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *BaseHiveParserListener) ExitAlterStatement(ctx *AlterStatementContext) {}

// EnterAlterTableStatementSuffix is called when production alterTableStatementSuffix is entered.
func (s *BaseHiveParserListener) EnterAlterTableStatementSuffix(ctx *AlterTableStatementSuffixContext) {
}

// ExitAlterTableStatementSuffix is called when production alterTableStatementSuffix is exited.
func (s *BaseHiveParserListener) ExitAlterTableStatementSuffix(ctx *AlterTableStatementSuffixContext) {
}

// EnterAlterTblPartitionStatementSuffix is called when production alterTblPartitionStatementSuffix is entered.
func (s *BaseHiveParserListener) EnterAlterTblPartitionStatementSuffix(ctx *AlterTblPartitionStatementSuffixContext) {
}

// ExitAlterTblPartitionStatementSuffix is called when production alterTblPartitionStatementSuffix is exited.
func (s *BaseHiveParserListener) ExitAlterTblPartitionStatementSuffix(ctx *AlterTblPartitionStatementSuffixContext) {
}

// EnterAlterStatementPartitionKeyType is called when production alterStatementPartitionKeyType is entered.
func (s *BaseHiveParserListener) EnterAlterStatementPartitionKeyType(ctx *AlterStatementPartitionKeyTypeContext) {
}

// ExitAlterStatementPartitionKeyType is called when production alterStatementPartitionKeyType is exited.
func (s *BaseHiveParserListener) ExitAlterStatementPartitionKeyType(ctx *AlterStatementPartitionKeyTypeContext) {
}

// EnterAlterViewStatementSuffix is called when production alterViewStatementSuffix is entered.
func (s *BaseHiveParserListener) EnterAlterViewStatementSuffix(ctx *AlterViewStatementSuffixContext) {
}

// ExitAlterViewStatementSuffix is called when production alterViewStatementSuffix is exited.
func (s *BaseHiveParserListener) ExitAlterViewStatementSuffix(ctx *AlterViewStatementSuffixContext) {}

// EnterAlterMaterializedViewStatementSuffix is called when production alterMaterializedViewStatementSuffix is entered.
func (s *BaseHiveParserListener) EnterAlterMaterializedViewStatementSuffix(ctx *AlterMaterializedViewStatementSuffixContext) {
}

// ExitAlterMaterializedViewStatementSuffix is called when production alterMaterializedViewStatementSuffix is exited.
func (s *BaseHiveParserListener) ExitAlterMaterializedViewStatementSuffix(ctx *AlterMaterializedViewStatementSuffixContext) {
}

// EnterAlterDatabaseStatementSuffix is called when production alterDatabaseStatementSuffix is entered.
func (s *BaseHiveParserListener) EnterAlterDatabaseStatementSuffix(ctx *AlterDatabaseStatementSuffixContext) {
}

// ExitAlterDatabaseStatementSuffix is called when production alterDatabaseStatementSuffix is exited.
func (s *BaseHiveParserListener) ExitAlterDatabaseStatementSuffix(ctx *AlterDatabaseStatementSuffixContext) {
}

// EnterAlterDatabaseSuffixProperties is called when production alterDatabaseSuffixProperties is entered.
func (s *BaseHiveParserListener) EnterAlterDatabaseSuffixProperties(ctx *AlterDatabaseSuffixPropertiesContext) {
}

// ExitAlterDatabaseSuffixProperties is called when production alterDatabaseSuffixProperties is exited.
func (s *BaseHiveParserListener) ExitAlterDatabaseSuffixProperties(ctx *AlterDatabaseSuffixPropertiesContext) {
}

// EnterAlterDatabaseSuffixSetOwner is called when production alterDatabaseSuffixSetOwner is entered.
func (s *BaseHiveParserListener) EnterAlterDatabaseSuffixSetOwner(ctx *AlterDatabaseSuffixSetOwnerContext) {
}

// ExitAlterDatabaseSuffixSetOwner is called when production alterDatabaseSuffixSetOwner is exited.
func (s *BaseHiveParserListener) ExitAlterDatabaseSuffixSetOwner(ctx *AlterDatabaseSuffixSetOwnerContext) {
}

// EnterAlterDatabaseSuffixSetLocation is called when production alterDatabaseSuffixSetLocation is entered.
func (s *BaseHiveParserListener) EnterAlterDatabaseSuffixSetLocation(ctx *AlterDatabaseSuffixSetLocationContext) {
}

// ExitAlterDatabaseSuffixSetLocation is called when production alterDatabaseSuffixSetLocation is exited.
func (s *BaseHiveParserListener) ExitAlterDatabaseSuffixSetLocation(ctx *AlterDatabaseSuffixSetLocationContext) {
}

// EnterAlterStatementSuffixRename is called when production alterStatementSuffixRename is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixRename(ctx *AlterStatementSuffixRenameContext) {
}

// ExitAlterStatementSuffixRename is called when production alterStatementSuffixRename is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixRename(ctx *AlterStatementSuffixRenameContext) {
}

// EnterAlterStatementSuffixAddCol is called when production alterStatementSuffixAddCol is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixAddCol(ctx *AlterStatementSuffixAddColContext) {
}

// ExitAlterStatementSuffixAddCol is called when production alterStatementSuffixAddCol is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixAddCol(ctx *AlterStatementSuffixAddColContext) {
}

// EnterAlterStatementSuffixAddConstraint is called when production alterStatementSuffixAddConstraint is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixAddConstraint(ctx *AlterStatementSuffixAddConstraintContext) {
}

// ExitAlterStatementSuffixAddConstraint is called when production alterStatementSuffixAddConstraint is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixAddConstraint(ctx *AlterStatementSuffixAddConstraintContext) {
}

// EnterAlterStatementSuffixUpdateColumns is called when production alterStatementSuffixUpdateColumns is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixUpdateColumns(ctx *AlterStatementSuffixUpdateColumnsContext) {
}

// ExitAlterStatementSuffixUpdateColumns is called when production alterStatementSuffixUpdateColumns is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixUpdateColumns(ctx *AlterStatementSuffixUpdateColumnsContext) {
}

// EnterAlterStatementSuffixDropConstraint is called when production alterStatementSuffixDropConstraint is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixDropConstraint(ctx *AlterStatementSuffixDropConstraintContext) {
}

// ExitAlterStatementSuffixDropConstraint is called when production alterStatementSuffixDropConstraint is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixDropConstraint(ctx *AlterStatementSuffixDropConstraintContext) {
}

// EnterAlterStatementSuffixRenameCol is called when production alterStatementSuffixRenameCol is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixRenameCol(ctx *AlterStatementSuffixRenameColContext) {
}

// ExitAlterStatementSuffixRenameCol is called when production alterStatementSuffixRenameCol is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixRenameCol(ctx *AlterStatementSuffixRenameColContext) {
}

// EnterAlterStatementSuffixUpdateStatsCol is called when production alterStatementSuffixUpdateStatsCol is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixUpdateStatsCol(ctx *AlterStatementSuffixUpdateStatsColContext) {
}

// ExitAlterStatementSuffixUpdateStatsCol is called when production alterStatementSuffixUpdateStatsCol is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixUpdateStatsCol(ctx *AlterStatementSuffixUpdateStatsColContext) {
}

// EnterAlterStatementSuffixUpdateStats is called when production alterStatementSuffixUpdateStats is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixUpdateStats(ctx *AlterStatementSuffixUpdateStatsContext) {
}

// ExitAlterStatementSuffixUpdateStats is called when production alterStatementSuffixUpdateStats is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixUpdateStats(ctx *AlterStatementSuffixUpdateStatsContext) {
}

// EnterAlterStatementChangeColPosition is called when production alterStatementChangeColPosition is entered.
func (s *BaseHiveParserListener) EnterAlterStatementChangeColPosition(ctx *AlterStatementChangeColPositionContext) {
}

// ExitAlterStatementChangeColPosition is called when production alterStatementChangeColPosition is exited.
func (s *BaseHiveParserListener) ExitAlterStatementChangeColPosition(ctx *AlterStatementChangeColPositionContext) {
}

// EnterAlterStatementSuffixAddPartitions is called when production alterStatementSuffixAddPartitions is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixAddPartitions(ctx *AlterStatementSuffixAddPartitionsContext) {
}

// ExitAlterStatementSuffixAddPartitions is called when production alterStatementSuffixAddPartitions is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixAddPartitions(ctx *AlterStatementSuffixAddPartitionsContext) {
}

// EnterAlterStatementSuffixAddPartitionsElement is called when production alterStatementSuffixAddPartitionsElement is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixAddPartitionsElement(ctx *AlterStatementSuffixAddPartitionsElementContext) {
}

// ExitAlterStatementSuffixAddPartitionsElement is called when production alterStatementSuffixAddPartitionsElement is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixAddPartitionsElement(ctx *AlterStatementSuffixAddPartitionsElementContext) {
}

// EnterAlterStatementSuffixTouch is called when production alterStatementSuffixTouch is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixTouch(ctx *AlterStatementSuffixTouchContext) {
}

// ExitAlterStatementSuffixTouch is called when production alterStatementSuffixTouch is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixTouch(ctx *AlterStatementSuffixTouchContext) {
}

// EnterAlterStatementSuffixArchive is called when production alterStatementSuffixArchive is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixArchive(ctx *AlterStatementSuffixArchiveContext) {
}

// ExitAlterStatementSuffixArchive is called when production alterStatementSuffixArchive is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixArchive(ctx *AlterStatementSuffixArchiveContext) {
}

// EnterAlterStatementSuffixUnArchive is called when production alterStatementSuffixUnArchive is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixUnArchive(ctx *AlterStatementSuffixUnArchiveContext) {
}

// ExitAlterStatementSuffixUnArchive is called when production alterStatementSuffixUnArchive is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixUnArchive(ctx *AlterStatementSuffixUnArchiveContext) {
}

// EnterPartitionLocation is called when production partitionLocation is entered.
func (s *BaseHiveParserListener) EnterPartitionLocation(ctx *PartitionLocationContext) {}

// ExitPartitionLocation is called when production partitionLocation is exited.
func (s *BaseHiveParserListener) ExitPartitionLocation(ctx *PartitionLocationContext) {}

// EnterAlterStatementSuffixDropPartitions is called when production alterStatementSuffixDropPartitions is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixDropPartitions(ctx *AlterStatementSuffixDropPartitionsContext) {
}

// ExitAlterStatementSuffixDropPartitions is called when production alterStatementSuffixDropPartitions is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixDropPartitions(ctx *AlterStatementSuffixDropPartitionsContext) {
}

// EnterAlterStatementSuffixProperties is called when production alterStatementSuffixProperties is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixProperties(ctx *AlterStatementSuffixPropertiesContext) {
}

// ExitAlterStatementSuffixProperties is called when production alterStatementSuffixProperties is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixProperties(ctx *AlterStatementSuffixPropertiesContext) {
}

// EnterAlterViewSuffixProperties is called when production alterViewSuffixProperties is entered.
func (s *BaseHiveParserListener) EnterAlterViewSuffixProperties(ctx *AlterViewSuffixPropertiesContext) {
}

// ExitAlterViewSuffixProperties is called when production alterViewSuffixProperties is exited.
func (s *BaseHiveParserListener) ExitAlterViewSuffixProperties(ctx *AlterViewSuffixPropertiesContext) {
}

// EnterAlterMaterializedViewSuffixRewrite is called when production alterMaterializedViewSuffixRewrite is entered.
func (s *BaseHiveParserListener) EnterAlterMaterializedViewSuffixRewrite(ctx *AlterMaterializedViewSuffixRewriteContext) {
}

// ExitAlterMaterializedViewSuffixRewrite is called when production alterMaterializedViewSuffixRewrite is exited.
func (s *BaseHiveParserListener) ExitAlterMaterializedViewSuffixRewrite(ctx *AlterMaterializedViewSuffixRewriteContext) {
}

// EnterAlterMaterializedViewSuffixRebuild is called when production alterMaterializedViewSuffixRebuild is entered.
func (s *BaseHiveParserListener) EnterAlterMaterializedViewSuffixRebuild(ctx *AlterMaterializedViewSuffixRebuildContext) {
}

// ExitAlterMaterializedViewSuffixRebuild is called when production alterMaterializedViewSuffixRebuild is exited.
func (s *BaseHiveParserListener) ExitAlterMaterializedViewSuffixRebuild(ctx *AlterMaterializedViewSuffixRebuildContext) {
}

// EnterAlterStatementSuffixSerdeProperties is called when production alterStatementSuffixSerdeProperties is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixSerdeProperties(ctx *AlterStatementSuffixSerdePropertiesContext) {
}

// ExitAlterStatementSuffixSerdeProperties is called when production alterStatementSuffixSerdeProperties is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixSerdeProperties(ctx *AlterStatementSuffixSerdePropertiesContext) {
}

// EnterTablePartitionPrefix is called when production tablePartitionPrefix is entered.
func (s *BaseHiveParserListener) EnterTablePartitionPrefix(ctx *TablePartitionPrefixContext) {}

// ExitTablePartitionPrefix is called when production tablePartitionPrefix is exited.
func (s *BaseHiveParserListener) ExitTablePartitionPrefix(ctx *TablePartitionPrefixContext) {}

// EnterAlterStatementSuffixFileFormat is called when production alterStatementSuffixFileFormat is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixFileFormat(ctx *AlterStatementSuffixFileFormatContext) {
}

// ExitAlterStatementSuffixFileFormat is called when production alterStatementSuffixFileFormat is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixFileFormat(ctx *AlterStatementSuffixFileFormatContext) {
}

// EnterAlterStatementSuffixClusterbySortby is called when production alterStatementSuffixClusterbySortby is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixClusterbySortby(ctx *AlterStatementSuffixClusterbySortbyContext) {
}

// ExitAlterStatementSuffixClusterbySortby is called when production alterStatementSuffixClusterbySortby is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixClusterbySortby(ctx *AlterStatementSuffixClusterbySortbyContext) {
}

// EnterAlterTblPartitionStatementSuffixSkewedLocation is called when production alterTblPartitionStatementSuffixSkewedLocation is entered.
func (s *BaseHiveParserListener) EnterAlterTblPartitionStatementSuffixSkewedLocation(ctx *AlterTblPartitionStatementSuffixSkewedLocationContext) {
}

// ExitAlterTblPartitionStatementSuffixSkewedLocation is called when production alterTblPartitionStatementSuffixSkewedLocation is exited.
func (s *BaseHiveParserListener) ExitAlterTblPartitionStatementSuffixSkewedLocation(ctx *AlterTblPartitionStatementSuffixSkewedLocationContext) {
}

// EnterSkewedLocations is called when production skewedLocations is entered.
func (s *BaseHiveParserListener) EnterSkewedLocations(ctx *SkewedLocationsContext) {}

// ExitSkewedLocations is called when production skewedLocations is exited.
func (s *BaseHiveParserListener) ExitSkewedLocations(ctx *SkewedLocationsContext) {}

// EnterSkewedLocationsList is called when production skewedLocationsList is entered.
func (s *BaseHiveParserListener) EnterSkewedLocationsList(ctx *SkewedLocationsListContext) {}

// ExitSkewedLocationsList is called when production skewedLocationsList is exited.
func (s *BaseHiveParserListener) ExitSkewedLocationsList(ctx *SkewedLocationsListContext) {}

// EnterSkewedLocationMap is called when production skewedLocationMap is entered.
func (s *BaseHiveParserListener) EnterSkewedLocationMap(ctx *SkewedLocationMapContext) {}

// ExitSkewedLocationMap is called when production skewedLocationMap is exited.
func (s *BaseHiveParserListener) ExitSkewedLocationMap(ctx *SkewedLocationMapContext) {}

// EnterAlterStatementSuffixLocation is called when production alterStatementSuffixLocation is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixLocation(ctx *AlterStatementSuffixLocationContext) {
}

// ExitAlterStatementSuffixLocation is called when production alterStatementSuffixLocation is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixLocation(ctx *AlterStatementSuffixLocationContext) {
}

// EnterAlterStatementSuffixSkewedby is called when production alterStatementSuffixSkewedby is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixSkewedby(ctx *AlterStatementSuffixSkewedbyContext) {
}

// ExitAlterStatementSuffixSkewedby is called when production alterStatementSuffixSkewedby is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixSkewedby(ctx *AlterStatementSuffixSkewedbyContext) {
}

// EnterAlterStatementSuffixExchangePartition is called when production alterStatementSuffixExchangePartition is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixExchangePartition(ctx *AlterStatementSuffixExchangePartitionContext) {
}

// ExitAlterStatementSuffixExchangePartition is called when production alterStatementSuffixExchangePartition is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixExchangePartition(ctx *AlterStatementSuffixExchangePartitionContext) {
}

// EnterAlterStatementSuffixRenamePart is called when production alterStatementSuffixRenamePart is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixRenamePart(ctx *AlterStatementSuffixRenamePartContext) {
}

// ExitAlterStatementSuffixRenamePart is called when production alterStatementSuffixRenamePart is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixRenamePart(ctx *AlterStatementSuffixRenamePartContext) {
}

// EnterAlterStatementSuffixStatsPart is called when production alterStatementSuffixStatsPart is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixStatsPart(ctx *AlterStatementSuffixStatsPartContext) {
}

// ExitAlterStatementSuffixStatsPart is called when production alterStatementSuffixStatsPart is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixStatsPart(ctx *AlterStatementSuffixStatsPartContext) {
}

// EnterAlterStatementSuffixMergeFiles is called when production alterStatementSuffixMergeFiles is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixMergeFiles(ctx *AlterStatementSuffixMergeFilesContext) {
}

// ExitAlterStatementSuffixMergeFiles is called when production alterStatementSuffixMergeFiles is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixMergeFiles(ctx *AlterStatementSuffixMergeFilesContext) {
}

// EnterAlterStatementSuffixBucketNum is called when production alterStatementSuffixBucketNum is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixBucketNum(ctx *AlterStatementSuffixBucketNumContext) {
}

// ExitAlterStatementSuffixBucketNum is called when production alterStatementSuffixBucketNum is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixBucketNum(ctx *AlterStatementSuffixBucketNumContext) {
}

// EnterBlocking is called when production blocking is entered.
func (s *BaseHiveParserListener) EnterBlocking(ctx *BlockingContext) {}

// ExitBlocking is called when production blocking is exited.
func (s *BaseHiveParserListener) ExitBlocking(ctx *BlockingContext) {}

// EnterAlterStatementSuffixCompact is called when production alterStatementSuffixCompact is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixCompact(ctx *AlterStatementSuffixCompactContext) {
}

// ExitAlterStatementSuffixCompact is called when production alterStatementSuffixCompact is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixCompact(ctx *AlterStatementSuffixCompactContext) {
}

// EnterAlterStatementSuffixSetOwner is called when production alterStatementSuffixSetOwner is entered.
func (s *BaseHiveParserListener) EnterAlterStatementSuffixSetOwner(ctx *AlterStatementSuffixSetOwnerContext) {
}

// ExitAlterStatementSuffixSetOwner is called when production alterStatementSuffixSetOwner is exited.
func (s *BaseHiveParserListener) ExitAlterStatementSuffixSetOwner(ctx *AlterStatementSuffixSetOwnerContext) {
}

// EnterFileFormat is called when production fileFormat is entered.
func (s *BaseHiveParserListener) EnterFileFormat(ctx *FileFormatContext) {}

// ExitFileFormat is called when production fileFormat is exited.
func (s *BaseHiveParserListener) ExitFileFormat(ctx *FileFormatContext) {}

// EnterInputFileFormat is called when production inputFileFormat is entered.
func (s *BaseHiveParserListener) EnterInputFileFormat(ctx *InputFileFormatContext) {}

// ExitInputFileFormat is called when production inputFileFormat is exited.
func (s *BaseHiveParserListener) ExitInputFileFormat(ctx *InputFileFormatContext) {}

// EnterTabTypeExpr is called when production tabTypeExpr is entered.
func (s *BaseHiveParserListener) EnterTabTypeExpr(ctx *TabTypeExprContext) {}

// ExitTabTypeExpr is called when production tabTypeExpr is exited.
func (s *BaseHiveParserListener) ExitTabTypeExpr(ctx *TabTypeExprContext) {}

// EnterPartTypeExpr is called when production partTypeExpr is entered.
func (s *BaseHiveParserListener) EnterPartTypeExpr(ctx *PartTypeExprContext) {}

// ExitPartTypeExpr is called when production partTypeExpr is exited.
func (s *BaseHiveParserListener) ExitPartTypeExpr(ctx *PartTypeExprContext) {}

// EnterTabPartColTypeExpr is called when production tabPartColTypeExpr is entered.
func (s *BaseHiveParserListener) EnterTabPartColTypeExpr(ctx *TabPartColTypeExprContext) {}

// ExitTabPartColTypeExpr is called when production tabPartColTypeExpr is exited.
func (s *BaseHiveParserListener) ExitTabPartColTypeExpr(ctx *TabPartColTypeExprContext) {}

// EnterDescStatement is called when production descStatement is entered.
func (s *BaseHiveParserListener) EnterDescStatement(ctx *DescStatementContext) {}

// ExitDescStatement is called when production descStatement is exited.
func (s *BaseHiveParserListener) ExitDescStatement(ctx *DescStatementContext) {}

// EnterAnalyzeStatement is called when production analyzeStatement is entered.
func (s *BaseHiveParserListener) EnterAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// ExitAnalyzeStatement is called when production analyzeStatement is exited.
func (s *BaseHiveParserListener) ExitAnalyzeStatement(ctx *AnalyzeStatementContext) {}

// EnterShowStatement is called when production showStatement is entered.
func (s *BaseHiveParserListener) EnterShowStatement(ctx *ShowStatementContext) {}

// ExitShowStatement is called when production showStatement is exited.
func (s *BaseHiveParserListener) ExitShowStatement(ctx *ShowStatementContext) {}

// EnterLockStatement is called when production lockStatement is entered.
func (s *BaseHiveParserListener) EnterLockStatement(ctx *LockStatementContext) {}

// ExitLockStatement is called when production lockStatement is exited.
func (s *BaseHiveParserListener) ExitLockStatement(ctx *LockStatementContext) {}

// EnterLockDatabase is called when production lockDatabase is entered.
func (s *BaseHiveParserListener) EnterLockDatabase(ctx *LockDatabaseContext) {}

// ExitLockDatabase is called when production lockDatabase is exited.
func (s *BaseHiveParserListener) ExitLockDatabase(ctx *LockDatabaseContext) {}

// EnterLockMode is called when production lockMode is entered.
func (s *BaseHiveParserListener) EnterLockMode(ctx *LockModeContext) {}

// ExitLockMode is called when production lockMode is exited.
func (s *BaseHiveParserListener) ExitLockMode(ctx *LockModeContext) {}

// EnterUnlockStatement is called when production unlockStatement is entered.
func (s *BaseHiveParserListener) EnterUnlockStatement(ctx *UnlockStatementContext) {}

// ExitUnlockStatement is called when production unlockStatement is exited.
func (s *BaseHiveParserListener) ExitUnlockStatement(ctx *UnlockStatementContext) {}

// EnterUnlockDatabase is called when production unlockDatabase is entered.
func (s *BaseHiveParserListener) EnterUnlockDatabase(ctx *UnlockDatabaseContext) {}

// ExitUnlockDatabase is called when production unlockDatabase is exited.
func (s *BaseHiveParserListener) ExitUnlockDatabase(ctx *UnlockDatabaseContext) {}

// EnterCreateRoleStatement is called when production createRoleStatement is entered.
func (s *BaseHiveParserListener) EnterCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// ExitCreateRoleStatement is called when production createRoleStatement is exited.
func (s *BaseHiveParserListener) ExitCreateRoleStatement(ctx *CreateRoleStatementContext) {}

// EnterDropRoleStatement is called when production dropRoleStatement is entered.
func (s *BaseHiveParserListener) EnterDropRoleStatement(ctx *DropRoleStatementContext) {}

// ExitDropRoleStatement is called when production dropRoleStatement is exited.
func (s *BaseHiveParserListener) ExitDropRoleStatement(ctx *DropRoleStatementContext) {}

// EnterGrantPrivileges is called when production grantPrivileges is entered.
func (s *BaseHiveParserListener) EnterGrantPrivileges(ctx *GrantPrivilegesContext) {}

// ExitGrantPrivileges is called when production grantPrivileges is exited.
func (s *BaseHiveParserListener) ExitGrantPrivileges(ctx *GrantPrivilegesContext) {}

// EnterRevokePrivileges is called when production revokePrivileges is entered.
func (s *BaseHiveParserListener) EnterRevokePrivileges(ctx *RevokePrivilegesContext) {}

// ExitRevokePrivileges is called when production revokePrivileges is exited.
func (s *BaseHiveParserListener) ExitRevokePrivileges(ctx *RevokePrivilegesContext) {}

// EnterGrantRole is called when production grantRole is entered.
func (s *BaseHiveParserListener) EnterGrantRole(ctx *GrantRoleContext) {}

// ExitGrantRole is called when production grantRole is exited.
func (s *BaseHiveParserListener) ExitGrantRole(ctx *GrantRoleContext) {}

// EnterRevokeRole is called when production revokeRole is entered.
func (s *BaseHiveParserListener) EnterRevokeRole(ctx *RevokeRoleContext) {}

// ExitRevokeRole is called when production revokeRole is exited.
func (s *BaseHiveParserListener) ExitRevokeRole(ctx *RevokeRoleContext) {}

// EnterShowRoleGrants is called when production showRoleGrants is entered.
func (s *BaseHiveParserListener) EnterShowRoleGrants(ctx *ShowRoleGrantsContext) {}

// ExitShowRoleGrants is called when production showRoleGrants is exited.
func (s *BaseHiveParserListener) ExitShowRoleGrants(ctx *ShowRoleGrantsContext) {}

// EnterShowRoles is called when production showRoles is entered.
func (s *BaseHiveParserListener) EnterShowRoles(ctx *ShowRolesContext) {}

// ExitShowRoles is called when production showRoles is exited.
func (s *BaseHiveParserListener) ExitShowRoles(ctx *ShowRolesContext) {}

// EnterShowCurrentRole is called when production showCurrentRole is entered.
func (s *BaseHiveParserListener) EnterShowCurrentRole(ctx *ShowCurrentRoleContext) {}

// ExitShowCurrentRole is called when production showCurrentRole is exited.
func (s *BaseHiveParserListener) ExitShowCurrentRole(ctx *ShowCurrentRoleContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *BaseHiveParserListener) EnterSetRole(ctx *SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *BaseHiveParserListener) ExitSetRole(ctx *SetRoleContext) {}

// EnterShowGrants is called when production showGrants is entered.
func (s *BaseHiveParserListener) EnterShowGrants(ctx *ShowGrantsContext) {}

// ExitShowGrants is called when production showGrants is exited.
func (s *BaseHiveParserListener) ExitShowGrants(ctx *ShowGrantsContext) {}

// EnterShowRolePrincipals is called when production showRolePrincipals is entered.
func (s *BaseHiveParserListener) EnterShowRolePrincipals(ctx *ShowRolePrincipalsContext) {}

// ExitShowRolePrincipals is called when production showRolePrincipals is exited.
func (s *BaseHiveParserListener) ExitShowRolePrincipals(ctx *ShowRolePrincipalsContext) {}

// EnterPrivilegeIncludeColObject is called when production privilegeIncludeColObject is entered.
func (s *BaseHiveParserListener) EnterPrivilegeIncludeColObject(ctx *PrivilegeIncludeColObjectContext) {
}

// ExitPrivilegeIncludeColObject is called when production privilegeIncludeColObject is exited.
func (s *BaseHiveParserListener) ExitPrivilegeIncludeColObject(ctx *PrivilegeIncludeColObjectContext) {
}

// EnterPrivilegeObject is called when production privilegeObject is entered.
func (s *BaseHiveParserListener) EnterPrivilegeObject(ctx *PrivilegeObjectContext) {}

// ExitPrivilegeObject is called when production privilegeObject is exited.
func (s *BaseHiveParserListener) ExitPrivilegeObject(ctx *PrivilegeObjectContext) {}

// EnterPrivObject is called when production privObject is entered.
func (s *BaseHiveParserListener) EnterPrivObject(ctx *PrivObjectContext) {}

// ExitPrivObject is called when production privObject is exited.
func (s *BaseHiveParserListener) ExitPrivObject(ctx *PrivObjectContext) {}

// EnterPrivObjectCols is called when production privObjectCols is entered.
func (s *BaseHiveParserListener) EnterPrivObjectCols(ctx *PrivObjectColsContext) {}

// ExitPrivObjectCols is called when production privObjectCols is exited.
func (s *BaseHiveParserListener) ExitPrivObjectCols(ctx *PrivObjectColsContext) {}

// EnterPrivilegeList is called when production privilegeList is entered.
func (s *BaseHiveParserListener) EnterPrivilegeList(ctx *PrivilegeListContext) {}

// ExitPrivilegeList is called when production privilegeList is exited.
func (s *BaseHiveParserListener) ExitPrivilegeList(ctx *PrivilegeListContext) {}

// EnterPrivlegeDef is called when production privlegeDef is entered.
func (s *BaseHiveParserListener) EnterPrivlegeDef(ctx *PrivlegeDefContext) {}

// ExitPrivlegeDef is called when production privlegeDef is exited.
func (s *BaseHiveParserListener) ExitPrivlegeDef(ctx *PrivlegeDefContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseHiveParserListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseHiveParserListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterPrincipalSpecification is called when production principalSpecification is entered.
func (s *BaseHiveParserListener) EnterPrincipalSpecification(ctx *PrincipalSpecificationContext) {}

// ExitPrincipalSpecification is called when production principalSpecification is exited.
func (s *BaseHiveParserListener) ExitPrincipalSpecification(ctx *PrincipalSpecificationContext) {}

// EnterPrincipalName is called when production principalName is entered.
func (s *BaseHiveParserListener) EnterPrincipalName(ctx *PrincipalNameContext) {}

// ExitPrincipalName is called when production principalName is exited.
func (s *BaseHiveParserListener) ExitPrincipalName(ctx *PrincipalNameContext) {}

// EnterWithGrantOption is called when production withGrantOption is entered.
func (s *BaseHiveParserListener) EnterWithGrantOption(ctx *WithGrantOptionContext) {}

// ExitWithGrantOption is called when production withGrantOption is exited.
func (s *BaseHiveParserListener) ExitWithGrantOption(ctx *WithGrantOptionContext) {}

// EnterGrantOptionFor is called when production grantOptionFor is entered.
func (s *BaseHiveParserListener) EnterGrantOptionFor(ctx *GrantOptionForContext) {}

// ExitGrantOptionFor is called when production grantOptionFor is exited.
func (s *BaseHiveParserListener) ExitGrantOptionFor(ctx *GrantOptionForContext) {}

// EnterAdminOptionFor is called when production adminOptionFor is entered.
func (s *BaseHiveParserListener) EnterAdminOptionFor(ctx *AdminOptionForContext) {}

// ExitAdminOptionFor is called when production adminOptionFor is exited.
func (s *BaseHiveParserListener) ExitAdminOptionFor(ctx *AdminOptionForContext) {}

// EnterWithAdminOption is called when production withAdminOption is entered.
func (s *BaseHiveParserListener) EnterWithAdminOption(ctx *WithAdminOptionContext) {}

// ExitWithAdminOption is called when production withAdminOption is exited.
func (s *BaseHiveParserListener) ExitWithAdminOption(ctx *WithAdminOptionContext) {}

// EnterMetastoreCheck is called when production metastoreCheck is entered.
func (s *BaseHiveParserListener) EnterMetastoreCheck(ctx *MetastoreCheckContext) {}

// ExitMetastoreCheck is called when production metastoreCheck is exited.
func (s *BaseHiveParserListener) ExitMetastoreCheck(ctx *MetastoreCheckContext) {}

// EnterResourceList is called when production resourceList is entered.
func (s *BaseHiveParserListener) EnterResourceList(ctx *ResourceListContext) {}

// ExitResourceList is called when production resourceList is exited.
func (s *BaseHiveParserListener) ExitResourceList(ctx *ResourceListContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseHiveParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseHiveParserListener) ExitResource(ctx *ResourceContext) {}

// EnterResourceType is called when production resourceType is entered.
func (s *BaseHiveParserListener) EnterResourceType(ctx *ResourceTypeContext) {}

// ExitResourceType is called when production resourceType is exited.
func (s *BaseHiveParserListener) ExitResourceType(ctx *ResourceTypeContext) {}

// EnterCreateFunctionStatement is called when production createFunctionStatement is entered.
func (s *BaseHiveParserListener) EnterCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// ExitCreateFunctionStatement is called when production createFunctionStatement is exited.
func (s *BaseHiveParserListener) ExitCreateFunctionStatement(ctx *CreateFunctionStatementContext) {}

// EnterDropFunctionStatement is called when production dropFunctionStatement is entered.
func (s *BaseHiveParserListener) EnterDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// ExitDropFunctionStatement is called when production dropFunctionStatement is exited.
func (s *BaseHiveParserListener) ExitDropFunctionStatement(ctx *DropFunctionStatementContext) {}

// EnterReloadFunctionStatement is called when production reloadFunctionStatement is entered.
func (s *BaseHiveParserListener) EnterReloadFunctionStatement(ctx *ReloadFunctionStatementContext) {}

// ExitReloadFunctionStatement is called when production reloadFunctionStatement is exited.
func (s *BaseHiveParserListener) ExitReloadFunctionStatement(ctx *ReloadFunctionStatementContext) {}

// EnterCreateMacroStatement is called when production createMacroStatement is entered.
func (s *BaseHiveParserListener) EnterCreateMacroStatement(ctx *CreateMacroStatementContext) {}

// ExitCreateMacroStatement is called when production createMacroStatement is exited.
func (s *BaseHiveParserListener) ExitCreateMacroStatement(ctx *CreateMacroStatementContext) {}

// EnterDropMacroStatement is called when production dropMacroStatement is entered.
func (s *BaseHiveParserListener) EnterDropMacroStatement(ctx *DropMacroStatementContext) {}

// ExitDropMacroStatement is called when production dropMacroStatement is exited.
func (s *BaseHiveParserListener) ExitDropMacroStatement(ctx *DropMacroStatementContext) {}

// EnterCreateViewStatement is called when production createViewStatement is entered.
func (s *BaseHiveParserListener) EnterCreateViewStatement(ctx *CreateViewStatementContext) {}

// ExitCreateViewStatement is called when production createViewStatement is exited.
func (s *BaseHiveParserListener) ExitCreateViewStatement(ctx *CreateViewStatementContext) {}

// EnterCreateMaterializedViewStatement is called when production createMaterializedViewStatement is entered.
func (s *BaseHiveParserListener) EnterCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// ExitCreateMaterializedViewStatement is called when production createMaterializedViewStatement is exited.
func (s *BaseHiveParserListener) ExitCreateMaterializedViewStatement(ctx *CreateMaterializedViewStatementContext) {
}

// EnterViewPartition is called when production viewPartition is entered.
func (s *BaseHiveParserListener) EnterViewPartition(ctx *ViewPartitionContext) {}

// ExitViewPartition is called when production viewPartition is exited.
func (s *BaseHiveParserListener) ExitViewPartition(ctx *ViewPartitionContext) {}

// EnterDropViewStatement is called when production dropViewStatement is entered.
func (s *BaseHiveParserListener) EnterDropViewStatement(ctx *DropViewStatementContext) {}

// ExitDropViewStatement is called when production dropViewStatement is exited.
func (s *BaseHiveParserListener) ExitDropViewStatement(ctx *DropViewStatementContext) {}

// EnterDropMaterializedViewStatement is called when production dropMaterializedViewStatement is entered.
func (s *BaseHiveParserListener) EnterDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// ExitDropMaterializedViewStatement is called when production dropMaterializedViewStatement is exited.
func (s *BaseHiveParserListener) ExitDropMaterializedViewStatement(ctx *DropMaterializedViewStatementContext) {
}

// EnterShowFunctionIdentifier is called when production showFunctionIdentifier is entered.
func (s *BaseHiveParserListener) EnterShowFunctionIdentifier(ctx *ShowFunctionIdentifierContext) {}

// ExitShowFunctionIdentifier is called when production showFunctionIdentifier is exited.
func (s *BaseHiveParserListener) ExitShowFunctionIdentifier(ctx *ShowFunctionIdentifierContext) {}

// EnterShowStmtIdentifier is called when production showStmtIdentifier is entered.
func (s *BaseHiveParserListener) EnterShowStmtIdentifier(ctx *ShowStmtIdentifierContext) {}

// ExitShowStmtIdentifier is called when production showStmtIdentifier is exited.
func (s *BaseHiveParserListener) ExitShowStmtIdentifier(ctx *ShowStmtIdentifierContext) {}

// EnterTableComment is called when production tableComment is entered.
func (s *BaseHiveParserListener) EnterTableComment(ctx *TableCommentContext) {}

// ExitTableComment is called when production tableComment is exited.
func (s *BaseHiveParserListener) ExitTableComment(ctx *TableCommentContext) {}

// EnterTablePartition is called when production tablePartition is entered.
func (s *BaseHiveParserListener) EnterTablePartition(ctx *TablePartitionContext) {}

// ExitTablePartition is called when production tablePartition is exited.
func (s *BaseHiveParserListener) ExitTablePartition(ctx *TablePartitionContext) {}

// EnterTableBuckets is called when production tableBuckets is entered.
func (s *BaseHiveParserListener) EnterTableBuckets(ctx *TableBucketsContext) {}

// ExitTableBuckets is called when production tableBuckets is exited.
func (s *BaseHiveParserListener) ExitTableBuckets(ctx *TableBucketsContext) {}

// EnterTableSkewed is called when production tableSkewed is entered.
func (s *BaseHiveParserListener) EnterTableSkewed(ctx *TableSkewedContext) {}

// ExitTableSkewed is called when production tableSkewed is exited.
func (s *BaseHiveParserListener) ExitTableSkewed(ctx *TableSkewedContext) {}

// EnterRowFormat is called when production rowFormat is entered.
func (s *BaseHiveParserListener) EnterRowFormat(ctx *RowFormatContext) {}

// ExitRowFormat is called when production rowFormat is exited.
func (s *BaseHiveParserListener) ExitRowFormat(ctx *RowFormatContext) {}

// EnterRecordReader is called when production recordReader is entered.
func (s *BaseHiveParserListener) EnterRecordReader(ctx *RecordReaderContext) {}

// ExitRecordReader is called when production recordReader is exited.
func (s *BaseHiveParserListener) ExitRecordReader(ctx *RecordReaderContext) {}

// EnterRecordWriter is called when production recordWriter is entered.
func (s *BaseHiveParserListener) EnterRecordWriter(ctx *RecordWriterContext) {}

// ExitRecordWriter is called when production recordWriter is exited.
func (s *BaseHiveParserListener) ExitRecordWriter(ctx *RecordWriterContext) {}

// EnterRowFormatSerde is called when production rowFormatSerde is entered.
func (s *BaseHiveParserListener) EnterRowFormatSerde(ctx *RowFormatSerdeContext) {}

// ExitRowFormatSerde is called when production rowFormatSerde is exited.
func (s *BaseHiveParserListener) ExitRowFormatSerde(ctx *RowFormatSerdeContext) {}

// EnterRowFormatDelimited is called when production rowFormatDelimited is entered.
func (s *BaseHiveParserListener) EnterRowFormatDelimited(ctx *RowFormatDelimitedContext) {}

// ExitRowFormatDelimited is called when production rowFormatDelimited is exited.
func (s *BaseHiveParserListener) ExitRowFormatDelimited(ctx *RowFormatDelimitedContext) {}

// EnterTableRowFormat is called when production tableRowFormat is entered.
func (s *BaseHiveParserListener) EnterTableRowFormat(ctx *TableRowFormatContext) {}

// ExitTableRowFormat is called when production tableRowFormat is exited.
func (s *BaseHiveParserListener) ExitTableRowFormat(ctx *TableRowFormatContext) {}

// EnterTablePropertiesPrefixed is called when production tablePropertiesPrefixed is entered.
func (s *BaseHiveParserListener) EnterTablePropertiesPrefixed(ctx *TablePropertiesPrefixedContext) {}

// ExitTablePropertiesPrefixed is called when production tablePropertiesPrefixed is exited.
func (s *BaseHiveParserListener) ExitTablePropertiesPrefixed(ctx *TablePropertiesPrefixedContext) {}

// EnterTableProperties is called when production tableProperties is entered.
func (s *BaseHiveParserListener) EnterTableProperties(ctx *TablePropertiesContext) {}

// ExitTableProperties is called when production tableProperties is exited.
func (s *BaseHiveParserListener) ExitTableProperties(ctx *TablePropertiesContext) {}

// EnterTablePropertiesList is called when production tablePropertiesList is entered.
func (s *BaseHiveParserListener) EnterTablePropertiesList(ctx *TablePropertiesListContext) {}

// ExitTablePropertiesList is called when production tablePropertiesList is exited.
func (s *BaseHiveParserListener) ExitTablePropertiesList(ctx *TablePropertiesListContext) {}

// EnterKeyValueProperty is called when production keyValueProperty is entered.
func (s *BaseHiveParserListener) EnterKeyValueProperty(ctx *KeyValuePropertyContext) {}

// ExitKeyValueProperty is called when production keyValueProperty is exited.
func (s *BaseHiveParserListener) ExitKeyValueProperty(ctx *KeyValuePropertyContext) {}

// EnterKeyProperty is called when production keyProperty is entered.
func (s *BaseHiveParserListener) EnterKeyProperty(ctx *KeyPropertyContext) {}

// ExitKeyProperty is called when production keyProperty is exited.
func (s *BaseHiveParserListener) ExitKeyProperty(ctx *KeyPropertyContext) {}

// EnterTableRowFormatFieldIdentifier is called when production tableRowFormatFieldIdentifier is entered.
func (s *BaseHiveParserListener) EnterTableRowFormatFieldIdentifier(ctx *TableRowFormatFieldIdentifierContext) {
}

// ExitTableRowFormatFieldIdentifier is called when production tableRowFormatFieldIdentifier is exited.
func (s *BaseHiveParserListener) ExitTableRowFormatFieldIdentifier(ctx *TableRowFormatFieldIdentifierContext) {
}

// EnterTableRowFormatCollItemsIdentifier is called when production tableRowFormatCollItemsIdentifier is entered.
func (s *BaseHiveParserListener) EnterTableRowFormatCollItemsIdentifier(ctx *TableRowFormatCollItemsIdentifierContext) {
}

// ExitTableRowFormatCollItemsIdentifier is called when production tableRowFormatCollItemsIdentifier is exited.
func (s *BaseHiveParserListener) ExitTableRowFormatCollItemsIdentifier(ctx *TableRowFormatCollItemsIdentifierContext) {
}

// EnterTableRowFormatMapKeysIdentifier is called when production tableRowFormatMapKeysIdentifier is entered.
func (s *BaseHiveParserListener) EnterTableRowFormatMapKeysIdentifier(ctx *TableRowFormatMapKeysIdentifierContext) {
}

// ExitTableRowFormatMapKeysIdentifier is called when production tableRowFormatMapKeysIdentifier is exited.
func (s *BaseHiveParserListener) ExitTableRowFormatMapKeysIdentifier(ctx *TableRowFormatMapKeysIdentifierContext) {
}

// EnterTableRowFormatLinesIdentifier is called when production tableRowFormatLinesIdentifier is entered.
func (s *BaseHiveParserListener) EnterTableRowFormatLinesIdentifier(ctx *TableRowFormatLinesIdentifierContext) {
}

// ExitTableRowFormatLinesIdentifier is called when production tableRowFormatLinesIdentifier is exited.
func (s *BaseHiveParserListener) ExitTableRowFormatLinesIdentifier(ctx *TableRowFormatLinesIdentifierContext) {
}

// EnterTableRowNullFormat is called when production tableRowNullFormat is entered.
func (s *BaseHiveParserListener) EnterTableRowNullFormat(ctx *TableRowNullFormatContext) {}

// ExitTableRowNullFormat is called when production tableRowNullFormat is exited.
func (s *BaseHiveParserListener) ExitTableRowNullFormat(ctx *TableRowNullFormatContext) {}

// EnterTableFileFormat is called when production tableFileFormat is entered.
func (s *BaseHiveParserListener) EnterTableFileFormat(ctx *TableFileFormatContext) {}

// ExitTableFileFormat is called when production tableFileFormat is exited.
func (s *BaseHiveParserListener) ExitTableFileFormat(ctx *TableFileFormatContext) {}

// EnterTableLocation is called when production tableLocation is entered.
func (s *BaseHiveParserListener) EnterTableLocation(ctx *TableLocationContext) {}

// ExitTableLocation is called when production tableLocation is exited.
func (s *BaseHiveParserListener) ExitTableLocation(ctx *TableLocationContext) {}

// EnterColumnNameTypeList is called when production columnNameTypeList is entered.
func (s *BaseHiveParserListener) EnterColumnNameTypeList(ctx *ColumnNameTypeListContext) {}

// ExitColumnNameTypeList is called when production columnNameTypeList is exited.
func (s *BaseHiveParserListener) ExitColumnNameTypeList(ctx *ColumnNameTypeListContext) {}

// EnterColumnNameTypeOrConstraintList is called when production columnNameTypeOrConstraintList is entered.
func (s *BaseHiveParserListener) EnterColumnNameTypeOrConstraintList(ctx *ColumnNameTypeOrConstraintListContext) {
}

// ExitColumnNameTypeOrConstraintList is called when production columnNameTypeOrConstraintList is exited.
func (s *BaseHiveParserListener) ExitColumnNameTypeOrConstraintList(ctx *ColumnNameTypeOrConstraintListContext) {
}

// EnterColumnNameColonTypeList is called when production columnNameColonTypeList is entered.
func (s *BaseHiveParserListener) EnterColumnNameColonTypeList(ctx *ColumnNameColonTypeListContext) {}

// ExitColumnNameColonTypeList is called when production columnNameColonTypeList is exited.
func (s *BaseHiveParserListener) ExitColumnNameColonTypeList(ctx *ColumnNameColonTypeListContext) {}

// EnterColumnNameList is called when production columnNameList is entered.
func (s *BaseHiveParserListener) EnterColumnNameList(ctx *ColumnNameListContext) {}

// ExitColumnNameList is called when production columnNameList is exited.
func (s *BaseHiveParserListener) ExitColumnNameList(ctx *ColumnNameListContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseHiveParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseHiveParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterExtColumnName is called when production extColumnName is entered.
func (s *BaseHiveParserListener) EnterExtColumnName(ctx *ExtColumnNameContext) {}

// ExitExtColumnName is called when production extColumnName is exited.
func (s *BaseHiveParserListener) ExitExtColumnName(ctx *ExtColumnNameContext) {}

// EnterColumnNameOrderList is called when production columnNameOrderList is entered.
func (s *BaseHiveParserListener) EnterColumnNameOrderList(ctx *ColumnNameOrderListContext) {}

// ExitColumnNameOrderList is called when production columnNameOrderList is exited.
func (s *BaseHiveParserListener) ExitColumnNameOrderList(ctx *ColumnNameOrderListContext) {}

// EnterColumnParenthesesList is called when production columnParenthesesList is entered.
func (s *BaseHiveParserListener) EnterColumnParenthesesList(ctx *ColumnParenthesesListContext) {}

// ExitColumnParenthesesList is called when production columnParenthesesList is exited.
func (s *BaseHiveParserListener) ExitColumnParenthesesList(ctx *ColumnParenthesesListContext) {}

// EnterEnableValidateSpecification is called when production enableValidateSpecification is entered.
func (s *BaseHiveParserListener) EnterEnableValidateSpecification(ctx *EnableValidateSpecificationContext) {
}

// ExitEnableValidateSpecification is called when production enableValidateSpecification is exited.
func (s *BaseHiveParserListener) ExitEnableValidateSpecification(ctx *EnableValidateSpecificationContext) {
}

// EnterEnableSpecification is called when production enableSpecification is entered.
func (s *BaseHiveParserListener) EnterEnableSpecification(ctx *EnableSpecificationContext) {}

// ExitEnableSpecification is called when production enableSpecification is exited.
func (s *BaseHiveParserListener) ExitEnableSpecification(ctx *EnableSpecificationContext) {}

// EnterValidateSpecification is called when production validateSpecification is entered.
func (s *BaseHiveParserListener) EnterValidateSpecification(ctx *ValidateSpecificationContext) {}

// ExitValidateSpecification is called when production validateSpecification is exited.
func (s *BaseHiveParserListener) ExitValidateSpecification(ctx *ValidateSpecificationContext) {}

// EnterEnforcedSpecification is called when production enforcedSpecification is entered.
func (s *BaseHiveParserListener) EnterEnforcedSpecification(ctx *EnforcedSpecificationContext) {}

// ExitEnforcedSpecification is called when production enforcedSpecification is exited.
func (s *BaseHiveParserListener) ExitEnforcedSpecification(ctx *EnforcedSpecificationContext) {}

// EnterRelySpecification is called when production relySpecification is entered.
func (s *BaseHiveParserListener) EnterRelySpecification(ctx *RelySpecificationContext) {}

// ExitRelySpecification is called when production relySpecification is exited.
func (s *BaseHiveParserListener) ExitRelySpecification(ctx *RelySpecificationContext) {}

// EnterCreateConstraint is called when production createConstraint is entered.
func (s *BaseHiveParserListener) EnterCreateConstraint(ctx *CreateConstraintContext) {}

// ExitCreateConstraint is called when production createConstraint is exited.
func (s *BaseHiveParserListener) ExitCreateConstraint(ctx *CreateConstraintContext) {}

// EnterAlterConstraintWithName is called when production alterConstraintWithName is entered.
func (s *BaseHiveParserListener) EnterAlterConstraintWithName(ctx *AlterConstraintWithNameContext) {}

// ExitAlterConstraintWithName is called when production alterConstraintWithName is exited.
func (s *BaseHiveParserListener) ExitAlterConstraintWithName(ctx *AlterConstraintWithNameContext) {}

// EnterTableLevelConstraint is called when production tableLevelConstraint is entered.
func (s *BaseHiveParserListener) EnterTableLevelConstraint(ctx *TableLevelConstraintContext) {}

// ExitTableLevelConstraint is called when production tableLevelConstraint is exited.
func (s *BaseHiveParserListener) ExitTableLevelConstraint(ctx *TableLevelConstraintContext) {}

// EnterPkUkConstraint is called when production pkUkConstraint is entered.
func (s *BaseHiveParserListener) EnterPkUkConstraint(ctx *PkUkConstraintContext) {}

// ExitPkUkConstraint is called when production pkUkConstraint is exited.
func (s *BaseHiveParserListener) ExitPkUkConstraint(ctx *PkUkConstraintContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseHiveParserListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseHiveParserListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterCreateForeignKey is called when production createForeignKey is entered.
func (s *BaseHiveParserListener) EnterCreateForeignKey(ctx *CreateForeignKeyContext) {}

// ExitCreateForeignKey is called when production createForeignKey is exited.
func (s *BaseHiveParserListener) ExitCreateForeignKey(ctx *CreateForeignKeyContext) {}

// EnterAlterForeignKeyWithName is called when production alterForeignKeyWithName is entered.
func (s *BaseHiveParserListener) EnterAlterForeignKeyWithName(ctx *AlterForeignKeyWithNameContext) {}

// ExitAlterForeignKeyWithName is called when production alterForeignKeyWithName is exited.
func (s *BaseHiveParserListener) ExitAlterForeignKeyWithName(ctx *AlterForeignKeyWithNameContext) {}

// EnterSkewedValueElement is called when production skewedValueElement is entered.
func (s *BaseHiveParserListener) EnterSkewedValueElement(ctx *SkewedValueElementContext) {}

// ExitSkewedValueElement is called when production skewedValueElement is exited.
func (s *BaseHiveParserListener) ExitSkewedValueElement(ctx *SkewedValueElementContext) {}

// EnterSkewedColumnValuePairList is called when production skewedColumnValuePairList is entered.
func (s *BaseHiveParserListener) EnterSkewedColumnValuePairList(ctx *SkewedColumnValuePairListContext) {
}

// ExitSkewedColumnValuePairList is called when production skewedColumnValuePairList is exited.
func (s *BaseHiveParserListener) ExitSkewedColumnValuePairList(ctx *SkewedColumnValuePairListContext) {
}

// EnterSkewedColumnValuePair is called when production skewedColumnValuePair is entered.
func (s *BaseHiveParserListener) EnterSkewedColumnValuePair(ctx *SkewedColumnValuePairContext) {}

// ExitSkewedColumnValuePair is called when production skewedColumnValuePair is exited.
func (s *BaseHiveParserListener) ExitSkewedColumnValuePair(ctx *SkewedColumnValuePairContext) {}

// EnterSkewedColumnValues is called when production skewedColumnValues is entered.
func (s *BaseHiveParserListener) EnterSkewedColumnValues(ctx *SkewedColumnValuesContext) {}

// ExitSkewedColumnValues is called when production skewedColumnValues is exited.
func (s *BaseHiveParserListener) ExitSkewedColumnValues(ctx *SkewedColumnValuesContext) {}

// EnterSkewedColumnValue is called when production skewedColumnValue is entered.
func (s *BaseHiveParserListener) EnterSkewedColumnValue(ctx *SkewedColumnValueContext) {}

// ExitSkewedColumnValue is called when production skewedColumnValue is exited.
func (s *BaseHiveParserListener) ExitSkewedColumnValue(ctx *SkewedColumnValueContext) {}

// EnterSkewedValueLocationElement is called when production skewedValueLocationElement is entered.
func (s *BaseHiveParserListener) EnterSkewedValueLocationElement(ctx *SkewedValueLocationElementContext) {
}

// ExitSkewedValueLocationElement is called when production skewedValueLocationElement is exited.
func (s *BaseHiveParserListener) ExitSkewedValueLocationElement(ctx *SkewedValueLocationElementContext) {
}

// EnterOrderSpecification is called when production orderSpecification is entered.
func (s *BaseHiveParserListener) EnterOrderSpecification(ctx *OrderSpecificationContext) {}

// ExitOrderSpecification is called when production orderSpecification is exited.
func (s *BaseHiveParserListener) ExitOrderSpecification(ctx *OrderSpecificationContext) {}

// EnterNullOrdering is called when production nullOrdering is entered.
func (s *BaseHiveParserListener) EnterNullOrdering(ctx *NullOrderingContext) {}

// ExitNullOrdering is called when production nullOrdering is exited.
func (s *BaseHiveParserListener) ExitNullOrdering(ctx *NullOrderingContext) {}

// EnterColumnNameOrder is called when production columnNameOrder is entered.
func (s *BaseHiveParserListener) EnterColumnNameOrder(ctx *ColumnNameOrderContext) {}

// ExitColumnNameOrder is called when production columnNameOrder is exited.
func (s *BaseHiveParserListener) ExitColumnNameOrder(ctx *ColumnNameOrderContext) {}

// EnterColumnNameCommentList is called when production columnNameCommentList is entered.
func (s *BaseHiveParserListener) EnterColumnNameCommentList(ctx *ColumnNameCommentListContext) {}

// ExitColumnNameCommentList is called when production columnNameCommentList is exited.
func (s *BaseHiveParserListener) ExitColumnNameCommentList(ctx *ColumnNameCommentListContext) {}

// EnterColumnNameComment is called when production columnNameComment is entered.
func (s *BaseHiveParserListener) EnterColumnNameComment(ctx *ColumnNameCommentContext) {}

// ExitColumnNameComment is called when production columnNameComment is exited.
func (s *BaseHiveParserListener) ExitColumnNameComment(ctx *ColumnNameCommentContext) {}

// EnterColumnRefOrder is called when production columnRefOrder is entered.
func (s *BaseHiveParserListener) EnterColumnRefOrder(ctx *ColumnRefOrderContext) {}

// ExitColumnRefOrder is called when production columnRefOrder is exited.
func (s *BaseHiveParserListener) ExitColumnRefOrder(ctx *ColumnRefOrderContext) {}

// EnterColumnNameType is called when production columnNameType is entered.
func (s *BaseHiveParserListener) EnterColumnNameType(ctx *ColumnNameTypeContext) {}

// ExitColumnNameType is called when production columnNameType is exited.
func (s *BaseHiveParserListener) ExitColumnNameType(ctx *ColumnNameTypeContext) {}

// EnterColumnNameTypeOrConstraint is called when production columnNameTypeOrConstraint is entered.
func (s *BaseHiveParserListener) EnterColumnNameTypeOrConstraint(ctx *ColumnNameTypeOrConstraintContext) {
}

// ExitColumnNameTypeOrConstraint is called when production columnNameTypeOrConstraint is exited.
func (s *BaseHiveParserListener) ExitColumnNameTypeOrConstraint(ctx *ColumnNameTypeOrConstraintContext) {
}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseHiveParserListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseHiveParserListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterColumnNameTypeConstraint is called when production columnNameTypeConstraint is entered.
func (s *BaseHiveParserListener) EnterColumnNameTypeConstraint(ctx *ColumnNameTypeConstraintContext) {
}

// ExitColumnNameTypeConstraint is called when production columnNameTypeConstraint is exited.
func (s *BaseHiveParserListener) ExitColumnNameTypeConstraint(ctx *ColumnNameTypeConstraintContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseHiveParserListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseHiveParserListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterForeignKeyConstraint is called when production foreignKeyConstraint is entered.
func (s *BaseHiveParserListener) EnterForeignKeyConstraint(ctx *ForeignKeyConstraintContext) {}

// ExitForeignKeyConstraint is called when production foreignKeyConstraint is exited.
func (s *BaseHiveParserListener) ExitForeignKeyConstraint(ctx *ForeignKeyConstraintContext) {}

// EnterColConstraint is called when production colConstraint is entered.
func (s *BaseHiveParserListener) EnterColConstraint(ctx *ColConstraintContext) {}

// ExitColConstraint is called when production colConstraint is exited.
func (s *BaseHiveParserListener) ExitColConstraint(ctx *ColConstraintContext) {}

// EnterAlterColumnConstraint is called when production alterColumnConstraint is entered.
func (s *BaseHiveParserListener) EnterAlterColumnConstraint(ctx *AlterColumnConstraintContext) {}

// ExitAlterColumnConstraint is called when production alterColumnConstraint is exited.
func (s *BaseHiveParserListener) ExitAlterColumnConstraint(ctx *AlterColumnConstraintContext) {}

// EnterAlterForeignKeyConstraint is called when production alterForeignKeyConstraint is entered.
func (s *BaseHiveParserListener) EnterAlterForeignKeyConstraint(ctx *AlterForeignKeyConstraintContext) {
}

// ExitAlterForeignKeyConstraint is called when production alterForeignKeyConstraint is exited.
func (s *BaseHiveParserListener) ExitAlterForeignKeyConstraint(ctx *AlterForeignKeyConstraintContext) {
}

// EnterAlterColConstraint is called when production alterColConstraint is entered.
func (s *BaseHiveParserListener) EnterAlterColConstraint(ctx *AlterColConstraintContext) {}

// ExitAlterColConstraint is called when production alterColConstraint is exited.
func (s *BaseHiveParserListener) ExitAlterColConstraint(ctx *AlterColConstraintContext) {}

// EnterColumnConstraintType is called when production columnConstraintType is entered.
func (s *BaseHiveParserListener) EnterColumnConstraintType(ctx *ColumnConstraintTypeContext) {}

// ExitColumnConstraintType is called when production columnConstraintType is exited.
func (s *BaseHiveParserListener) ExitColumnConstraintType(ctx *ColumnConstraintTypeContext) {}

// EnterDefaultVal is called when production defaultVal is entered.
func (s *BaseHiveParserListener) EnterDefaultVal(ctx *DefaultValContext) {}

// ExitDefaultVal is called when production defaultVal is exited.
func (s *BaseHiveParserListener) ExitDefaultVal(ctx *DefaultValContext) {}

// EnterTableConstraintType is called when production tableConstraintType is entered.
func (s *BaseHiveParserListener) EnterTableConstraintType(ctx *TableConstraintTypeContext) {}

// ExitTableConstraintType is called when production tableConstraintType is exited.
func (s *BaseHiveParserListener) ExitTableConstraintType(ctx *TableConstraintTypeContext) {}

// EnterConstraintOptsCreate is called when production constraintOptsCreate is entered.
func (s *BaseHiveParserListener) EnterConstraintOptsCreate(ctx *ConstraintOptsCreateContext) {}

// ExitConstraintOptsCreate is called when production constraintOptsCreate is exited.
func (s *BaseHiveParserListener) ExitConstraintOptsCreate(ctx *ConstraintOptsCreateContext) {}

// EnterConstraintOptsAlter is called when production constraintOptsAlter is entered.
func (s *BaseHiveParserListener) EnterConstraintOptsAlter(ctx *ConstraintOptsAlterContext) {}

// ExitConstraintOptsAlter is called when production constraintOptsAlter is exited.
func (s *BaseHiveParserListener) ExitConstraintOptsAlter(ctx *ConstraintOptsAlterContext) {}

// EnterColumnNameColonType is called when production columnNameColonType is entered.
func (s *BaseHiveParserListener) EnterColumnNameColonType(ctx *ColumnNameColonTypeContext) {}

// ExitColumnNameColonType is called when production columnNameColonType is exited.
func (s *BaseHiveParserListener) ExitColumnNameColonType(ctx *ColumnNameColonTypeContext) {}

// EnterColType is called when production colType is entered.
func (s *BaseHiveParserListener) EnterColType(ctx *ColTypeContext) {}

// ExitColType is called when production colType is exited.
func (s *BaseHiveParserListener) ExitColType(ctx *ColTypeContext) {}

// EnterColTypeList is called when production colTypeList is entered.
func (s *BaseHiveParserListener) EnterColTypeList(ctx *ColTypeListContext) {}

// ExitColTypeList is called when production colTypeList is exited.
func (s *BaseHiveParserListener) ExitColTypeList(ctx *ColTypeListContext) {}

// EnterType_db_col is called when production type_db_col is entered.
func (s *BaseHiveParserListener) EnterType_db_col(ctx *Type_db_colContext) {}

// ExitType_db_col is called when production type_db_col is exited.
func (s *BaseHiveParserListener) ExitType_db_col(ctx *Type_db_colContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseHiveParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseHiveParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseHiveParserListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseHiveParserListener) ExitListType(ctx *ListTypeContext) {}

// EnterStructType is called when production structType is entered.
func (s *BaseHiveParserListener) EnterStructType(ctx *StructTypeContext) {}

// ExitStructType is called when production structType is exited.
func (s *BaseHiveParserListener) ExitStructType(ctx *StructTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseHiveParserListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseHiveParserListener) ExitMapType(ctx *MapTypeContext) {}

// EnterUnionType is called when production unionType is entered.
func (s *BaseHiveParserListener) EnterUnionType(ctx *UnionTypeContext) {}

// ExitUnionType is called when production unionType is exited.
func (s *BaseHiveParserListener) ExitUnionType(ctx *UnionTypeContext) {}

// EnterSetOperator is called when production setOperator is entered.
func (s *BaseHiveParserListener) EnterSetOperator(ctx *SetOperatorContext) {}

// ExitSetOperator is called when production setOperator is exited.
func (s *BaseHiveParserListener) ExitSetOperator(ctx *SetOperatorContext) {}

// EnterQueryStatementExpression is called when production queryStatementExpression is entered.
func (s *BaseHiveParserListener) EnterQueryStatementExpression(ctx *QueryStatementExpressionContext) {
}

// ExitQueryStatementExpression is called when production queryStatementExpression is exited.
func (s *BaseHiveParserListener) ExitQueryStatementExpression(ctx *QueryStatementExpressionContext) {}

// EnterQueryStatementExpressionBody is called when production queryStatementExpressionBody is entered.
func (s *BaseHiveParserListener) EnterQueryStatementExpressionBody(ctx *QueryStatementExpressionBodyContext) {
}

// ExitQueryStatementExpressionBody is called when production queryStatementExpressionBody is exited.
func (s *BaseHiveParserListener) ExitQueryStatementExpressionBody(ctx *QueryStatementExpressionBodyContext) {
}

// EnterWithClause is called when production withClause is entered.
func (s *BaseHiveParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseHiveParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterCteStatement is called when production cteStatement is entered.
func (s *BaseHiveParserListener) EnterCteStatement(ctx *CteStatementContext) {}

// ExitCteStatement is called when production cteStatement is exited.
func (s *BaseHiveParserListener) ExitCteStatement(ctx *CteStatementContext) {}

// EnterFromStatement is called when production fromStatement is entered.
func (s *BaseHiveParserListener) EnterFromStatement(ctx *FromStatementContext) {}

// ExitFromStatement is called when production fromStatement is exited.
func (s *BaseHiveParserListener) ExitFromStatement(ctx *FromStatementContext) {}

// EnterSingleFromStatement is called when production singleFromStatement is entered.
func (s *BaseHiveParserListener) EnterSingleFromStatement(ctx *SingleFromStatementContext) {}

// ExitSingleFromStatement is called when production singleFromStatement is exited.
func (s *BaseHiveParserListener) ExitSingleFromStatement(ctx *SingleFromStatementContext) {}

// EnterRegularBody is called when production regularBody is entered.
func (s *BaseHiveParserListener) EnterRegularBody(ctx *RegularBodyContext) {}

// ExitRegularBody is called when production regularBody is exited.
func (s *BaseHiveParserListener) ExitRegularBody(ctx *RegularBodyContext) {}

// EnterAtomSelectStatement is called when production atomSelectStatement is entered.
func (s *BaseHiveParserListener) EnterAtomSelectStatement(ctx *AtomSelectStatementContext) {}

// ExitAtomSelectStatement is called when production atomSelectStatement is exited.
func (s *BaseHiveParserListener) ExitAtomSelectStatement(ctx *AtomSelectStatementContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseHiveParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseHiveParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterSetOpSelectStatement is called when production setOpSelectStatement is entered.
func (s *BaseHiveParserListener) EnterSetOpSelectStatement(ctx *SetOpSelectStatementContext) {}

// ExitSetOpSelectStatement is called when production setOpSelectStatement is exited.
func (s *BaseHiveParserListener) ExitSetOpSelectStatement(ctx *SetOpSelectStatementContext) {}

// EnterSelectStatementWithCTE is called when production selectStatementWithCTE is entered.
func (s *BaseHiveParserListener) EnterSelectStatementWithCTE(ctx *SelectStatementWithCTEContext) {}

// ExitSelectStatementWithCTE is called when production selectStatementWithCTE is exited.
func (s *BaseHiveParserListener) ExitSelectStatementWithCTE(ctx *SelectStatementWithCTEContext) {}

// EnterBody is called when production body is entered.
func (s *BaseHiveParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseHiveParserListener) ExitBody(ctx *BodyContext) {}

// EnterInsertClause is called when production insertClause is entered.
func (s *BaseHiveParserListener) EnterInsertClause(ctx *InsertClauseContext) {}

// ExitInsertClause is called when production insertClause is exited.
func (s *BaseHiveParserListener) ExitInsertClause(ctx *InsertClauseContext) {}

// EnterDestination is called when production destination is entered.
func (s *BaseHiveParserListener) EnterDestination(ctx *DestinationContext) {}

// ExitDestination is called when production destination is exited.
func (s *BaseHiveParserListener) ExitDestination(ctx *DestinationContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseHiveParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseHiveParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseHiveParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseHiveParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterColumnAssignmentClause is called when production columnAssignmentClause is entered.
func (s *BaseHiveParserListener) EnterColumnAssignmentClause(ctx *ColumnAssignmentClauseContext) {}

// ExitColumnAssignmentClause is called when production columnAssignmentClause is exited.
func (s *BaseHiveParserListener) ExitColumnAssignmentClause(ctx *ColumnAssignmentClauseContext) {}

// EnterSetColumnsClause is called when production setColumnsClause is entered.
func (s *BaseHiveParserListener) EnterSetColumnsClause(ctx *SetColumnsClauseContext) {}

// ExitSetColumnsClause is called when production setColumnsClause is exited.
func (s *BaseHiveParserListener) ExitSetColumnsClause(ctx *SetColumnsClauseContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseHiveParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseHiveParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterSqlTransactionStatement is called when production sqlTransactionStatement is entered.
func (s *BaseHiveParserListener) EnterSqlTransactionStatement(ctx *SqlTransactionStatementContext) {}

// ExitSqlTransactionStatement is called when production sqlTransactionStatement is exited.
func (s *BaseHiveParserListener) ExitSqlTransactionStatement(ctx *SqlTransactionStatementContext) {}

// EnterStartTransactionStatement is called when production startTransactionStatement is entered.
func (s *BaseHiveParserListener) EnterStartTransactionStatement(ctx *StartTransactionStatementContext) {
}

// ExitStartTransactionStatement is called when production startTransactionStatement is exited.
func (s *BaseHiveParserListener) ExitStartTransactionStatement(ctx *StartTransactionStatementContext) {
}

// EnterTransactionMode is called when production transactionMode is entered.
func (s *BaseHiveParserListener) EnterTransactionMode(ctx *TransactionModeContext) {}

// ExitTransactionMode is called when production transactionMode is exited.
func (s *BaseHiveParserListener) ExitTransactionMode(ctx *TransactionModeContext) {}

// EnterTransactionAccessMode is called when production transactionAccessMode is entered.
func (s *BaseHiveParserListener) EnterTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// ExitTransactionAccessMode is called when production transactionAccessMode is exited.
func (s *BaseHiveParserListener) ExitTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// EnterIsolationLevel is called when production isolationLevel is entered.
func (s *BaseHiveParserListener) EnterIsolationLevel(ctx *IsolationLevelContext) {}

// ExitIsolationLevel is called when production isolationLevel is exited.
func (s *BaseHiveParserListener) ExitIsolationLevel(ctx *IsolationLevelContext) {}

// EnterLevelOfIsolation is called when production levelOfIsolation is entered.
func (s *BaseHiveParserListener) EnterLevelOfIsolation(ctx *LevelOfIsolationContext) {}

// ExitLevelOfIsolation is called when production levelOfIsolation is exited.
func (s *BaseHiveParserListener) ExitLevelOfIsolation(ctx *LevelOfIsolationContext) {}

// EnterCommitStatement is called when production commitStatement is entered.
func (s *BaseHiveParserListener) EnterCommitStatement(ctx *CommitStatementContext) {}

// ExitCommitStatement is called when production commitStatement is exited.
func (s *BaseHiveParserListener) ExitCommitStatement(ctx *CommitStatementContext) {}

// EnterRollbackStatement is called when production rollbackStatement is entered.
func (s *BaseHiveParserListener) EnterRollbackStatement(ctx *RollbackStatementContext) {}

// ExitRollbackStatement is called when production rollbackStatement is exited.
func (s *BaseHiveParserListener) ExitRollbackStatement(ctx *RollbackStatementContext) {}

// EnterSetAutoCommitStatement is called when production setAutoCommitStatement is entered.
func (s *BaseHiveParserListener) EnterSetAutoCommitStatement(ctx *SetAutoCommitStatementContext) {}

// ExitSetAutoCommitStatement is called when production setAutoCommitStatement is exited.
func (s *BaseHiveParserListener) ExitSetAutoCommitStatement(ctx *SetAutoCommitStatementContext) {}

// EnterAbortTransactionStatement is called when production abortTransactionStatement is entered.
func (s *BaseHiveParserListener) EnterAbortTransactionStatement(ctx *AbortTransactionStatementContext) {
}

// ExitAbortTransactionStatement is called when production abortTransactionStatement is exited.
func (s *BaseHiveParserListener) ExitAbortTransactionStatement(ctx *AbortTransactionStatementContext) {
}

// EnterMergeStatement is called when production mergeStatement is entered.
func (s *BaseHiveParserListener) EnterMergeStatement(ctx *MergeStatementContext) {}

// ExitMergeStatement is called when production mergeStatement is exited.
func (s *BaseHiveParserListener) ExitMergeStatement(ctx *MergeStatementContext) {}

// EnterWhenClauses is called when production whenClauses is entered.
func (s *BaseHiveParserListener) EnterWhenClauses(ctx *WhenClausesContext) {}

// ExitWhenClauses is called when production whenClauses is exited.
func (s *BaseHiveParserListener) ExitWhenClauses(ctx *WhenClausesContext) {}

// EnterWhenNotMatchedClause is called when production whenNotMatchedClause is entered.
func (s *BaseHiveParserListener) EnterWhenNotMatchedClause(ctx *WhenNotMatchedClauseContext) {}

// ExitWhenNotMatchedClause is called when production whenNotMatchedClause is exited.
func (s *BaseHiveParserListener) ExitWhenNotMatchedClause(ctx *WhenNotMatchedClauseContext) {}

// EnterWhenMatchedAndClause is called when production whenMatchedAndClause is entered.
func (s *BaseHiveParserListener) EnterWhenMatchedAndClause(ctx *WhenMatchedAndClauseContext) {}

// ExitWhenMatchedAndClause is called when production whenMatchedAndClause is exited.
func (s *BaseHiveParserListener) ExitWhenMatchedAndClause(ctx *WhenMatchedAndClauseContext) {}

// EnterWhenMatchedThenClause is called when production whenMatchedThenClause is entered.
func (s *BaseHiveParserListener) EnterWhenMatchedThenClause(ctx *WhenMatchedThenClauseContext) {}

// ExitWhenMatchedThenClause is called when production whenMatchedThenClause is exited.
func (s *BaseHiveParserListener) ExitWhenMatchedThenClause(ctx *WhenMatchedThenClauseContext) {}

// EnterUpdateOrDelete is called when production updateOrDelete is entered.
func (s *BaseHiveParserListener) EnterUpdateOrDelete(ctx *UpdateOrDeleteContext) {}

// ExitUpdateOrDelete is called when production updateOrDelete is exited.
func (s *BaseHiveParserListener) ExitUpdateOrDelete(ctx *UpdateOrDeleteContext) {}

// EnterKillQueryStatement is called when production killQueryStatement is entered.
func (s *BaseHiveParserListener) EnterKillQueryStatement(ctx *KillQueryStatementContext) {}

// ExitKillQueryStatement is called when production killQueryStatement is exited.
func (s *BaseHiveParserListener) ExitKillQueryStatement(ctx *KillQueryStatementContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseHiveParserListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseHiveParserListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterSelectList is called when production selectList is entered.
func (s *BaseHiveParserListener) EnterSelectList(ctx *SelectListContext) {}

// ExitSelectList is called when production selectList is exited.
func (s *BaseHiveParserListener) ExitSelectList(ctx *SelectListContext) {}

// EnterSelectTrfmClause is called when production selectTrfmClause is entered.
func (s *BaseHiveParserListener) EnterSelectTrfmClause(ctx *SelectTrfmClauseContext) {}

// ExitSelectTrfmClause is called when production selectTrfmClause is exited.
func (s *BaseHiveParserListener) ExitSelectTrfmClause(ctx *SelectTrfmClauseContext) {}

// EnterSelectItem is called when production selectItem is entered.
func (s *BaseHiveParserListener) EnterSelectItem(ctx *SelectItemContext) {}

// ExitSelectItem is called when production selectItem is exited.
func (s *BaseHiveParserListener) ExitSelectItem(ctx *SelectItemContext) {}

// EnterTrfmClause is called when production trfmClause is entered.
func (s *BaseHiveParserListener) EnterTrfmClause(ctx *TrfmClauseContext) {}

// ExitTrfmClause is called when production trfmClause is exited.
func (s *BaseHiveParserListener) ExitTrfmClause(ctx *TrfmClauseContext) {}

// EnterSelectExpression is called when production selectExpression is entered.
func (s *BaseHiveParserListener) EnterSelectExpression(ctx *SelectExpressionContext) {}

// ExitSelectExpression is called when production selectExpression is exited.
func (s *BaseHiveParserListener) ExitSelectExpression(ctx *SelectExpressionContext) {}

// EnterSelectExpressionList is called when production selectExpressionList is entered.
func (s *BaseHiveParserListener) EnterSelectExpressionList(ctx *SelectExpressionListContext) {}

// ExitSelectExpressionList is called when production selectExpressionList is exited.
func (s *BaseHiveParserListener) ExitSelectExpressionList(ctx *SelectExpressionListContext) {}

// EnterWindow_clause is called when production window_clause is entered.
func (s *BaseHiveParserListener) EnterWindow_clause(ctx *Window_clauseContext) {}

// ExitWindow_clause is called when production window_clause is exited.
func (s *BaseHiveParserListener) ExitWindow_clause(ctx *Window_clauseContext) {}

// EnterWindow_defn is called when production window_defn is entered.
func (s *BaseHiveParserListener) EnterWindow_defn(ctx *Window_defnContext) {}

// ExitWindow_defn is called when production window_defn is exited.
func (s *BaseHiveParserListener) ExitWindow_defn(ctx *Window_defnContext) {}

// EnterWindow_specification is called when production window_specification is entered.
func (s *BaseHiveParserListener) EnterWindow_specification(ctx *Window_specificationContext) {}

// ExitWindow_specification is called when production window_specification is exited.
func (s *BaseHiveParserListener) ExitWindow_specification(ctx *Window_specificationContext) {}

// EnterWindow_frame is called when production window_frame is entered.
func (s *BaseHiveParserListener) EnterWindow_frame(ctx *Window_frameContext) {}

// ExitWindow_frame is called when production window_frame is exited.
func (s *BaseHiveParserListener) ExitWindow_frame(ctx *Window_frameContext) {}

// EnterWindow_range_expression is called when production window_range_expression is entered.
func (s *BaseHiveParserListener) EnterWindow_range_expression(ctx *Window_range_expressionContext) {}

// ExitWindow_range_expression is called when production window_range_expression is exited.
func (s *BaseHiveParserListener) ExitWindow_range_expression(ctx *Window_range_expressionContext) {}

// EnterWindow_value_expression is called when production window_value_expression is entered.
func (s *BaseHiveParserListener) EnterWindow_value_expression(ctx *Window_value_expressionContext) {}

// ExitWindow_value_expression is called when production window_value_expression is exited.
func (s *BaseHiveParserListener) ExitWindow_value_expression(ctx *Window_value_expressionContext) {}

// EnterWindow_frame_start_boundary is called when production window_frame_start_boundary is entered.
func (s *BaseHiveParserListener) EnterWindow_frame_start_boundary(ctx *Window_frame_start_boundaryContext) {
}

// ExitWindow_frame_start_boundary is called when production window_frame_start_boundary is exited.
func (s *BaseHiveParserListener) ExitWindow_frame_start_boundary(ctx *Window_frame_start_boundaryContext) {
}

// EnterWindow_frame_boundary is called when production window_frame_boundary is entered.
func (s *BaseHiveParserListener) EnterWindow_frame_boundary(ctx *Window_frame_boundaryContext) {}

// ExitWindow_frame_boundary is called when production window_frame_boundary is exited.
func (s *BaseHiveParserListener) ExitWindow_frame_boundary(ctx *Window_frame_boundaryContext) {}

// EnterTableAllColumns is called when production tableAllColumns is entered.
func (s *BaseHiveParserListener) EnterTableAllColumns(ctx *TableAllColumnsContext) {}

// ExitTableAllColumns is called when production tableAllColumns is exited.
func (s *BaseHiveParserListener) ExitTableAllColumns(ctx *TableAllColumnsContext) {}

// EnterTableOrColumn is called when production tableOrColumn is entered.
func (s *BaseHiveParserListener) EnterTableOrColumn(ctx *TableOrColumnContext) {}

// ExitTableOrColumn is called when production tableOrColumn is exited.
func (s *BaseHiveParserListener) ExitTableOrColumn(ctx *TableOrColumnContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseHiveParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseHiveParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterAliasList is called when production aliasList is entered.
func (s *BaseHiveParserListener) EnterAliasList(ctx *AliasListContext) {}

// ExitAliasList is called when production aliasList is exited.
func (s *BaseHiveParserListener) ExitAliasList(ctx *AliasListContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseHiveParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseHiveParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterFromSource is called when production fromSource is entered.
func (s *BaseHiveParserListener) EnterFromSource(ctx *FromSourceContext) {}

// ExitFromSource is called when production fromSource is exited.
func (s *BaseHiveParserListener) ExitFromSource(ctx *FromSourceContext) {}

// EnterAtomjoinSource is called when production atomjoinSource is entered.
func (s *BaseHiveParserListener) EnterAtomjoinSource(ctx *AtomjoinSourceContext) {}

// ExitAtomjoinSource is called when production atomjoinSource is exited.
func (s *BaseHiveParserListener) ExitAtomjoinSource(ctx *AtomjoinSourceContext) {}

// EnterJoinSource is called when production joinSource is entered.
func (s *BaseHiveParserListener) EnterJoinSource(ctx *JoinSourceContext) {}

// ExitJoinSource is called when production joinSource is exited.
func (s *BaseHiveParserListener) ExitJoinSource(ctx *JoinSourceContext) {}

// EnterJoinSourcePart is called when production joinSourcePart is entered.
func (s *BaseHiveParserListener) EnterJoinSourcePart(ctx *JoinSourcePartContext) {}

// ExitJoinSourcePart is called when production joinSourcePart is exited.
func (s *BaseHiveParserListener) ExitJoinSourcePart(ctx *JoinSourcePartContext) {}

// EnterUniqueJoinSource is called when production uniqueJoinSource is entered.
func (s *BaseHiveParserListener) EnterUniqueJoinSource(ctx *UniqueJoinSourceContext) {}

// ExitUniqueJoinSource is called when production uniqueJoinSource is exited.
func (s *BaseHiveParserListener) ExitUniqueJoinSource(ctx *UniqueJoinSourceContext) {}

// EnterUniqueJoinExpr is called when production uniqueJoinExpr is entered.
func (s *BaseHiveParserListener) EnterUniqueJoinExpr(ctx *UniqueJoinExprContext) {}

// ExitUniqueJoinExpr is called when production uniqueJoinExpr is exited.
func (s *BaseHiveParserListener) ExitUniqueJoinExpr(ctx *UniqueJoinExprContext) {}

// EnterUniqueJoinToken is called when production uniqueJoinToken is entered.
func (s *BaseHiveParserListener) EnterUniqueJoinToken(ctx *UniqueJoinTokenContext) {}

// ExitUniqueJoinToken is called when production uniqueJoinToken is exited.
func (s *BaseHiveParserListener) ExitUniqueJoinToken(ctx *UniqueJoinTokenContext) {}

// EnterJoinToken is called when production joinToken is entered.
func (s *BaseHiveParserListener) EnterJoinToken(ctx *JoinTokenContext) {}

// ExitJoinToken is called when production joinToken is exited.
func (s *BaseHiveParserListener) ExitJoinToken(ctx *JoinTokenContext) {}

// EnterLateralView is called when production lateralView is entered.
func (s *BaseHiveParserListener) EnterLateralView(ctx *LateralViewContext) {}

// ExitLateralView is called when production lateralView is exited.
func (s *BaseHiveParserListener) ExitLateralView(ctx *LateralViewContext) {}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseHiveParserListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseHiveParserListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterTableBucketSample is called when production tableBucketSample is entered.
func (s *BaseHiveParserListener) EnterTableBucketSample(ctx *TableBucketSampleContext) {}

// ExitTableBucketSample is called when production tableBucketSample is exited.
func (s *BaseHiveParserListener) ExitTableBucketSample(ctx *TableBucketSampleContext) {}

// EnterSplitSample is called when production splitSample is entered.
func (s *BaseHiveParserListener) EnterSplitSample(ctx *SplitSampleContext) {}

// ExitSplitSample is called when production splitSample is exited.
func (s *BaseHiveParserListener) ExitSplitSample(ctx *SplitSampleContext) {}

// EnterTableSample is called when production tableSample is entered.
func (s *BaseHiveParserListener) EnterTableSample(ctx *TableSampleContext) {}

// ExitTableSample is called when production tableSample is exited.
func (s *BaseHiveParserListener) ExitTableSample(ctx *TableSampleContext) {}

// EnterTableSource is called when production tableSource is entered.
func (s *BaseHiveParserListener) EnterTableSource(ctx *TableSourceContext) {}

// ExitTableSource is called when production tableSource is exited.
func (s *BaseHiveParserListener) ExitTableSource(ctx *TableSourceContext) {}

// EnterUniqueJoinTableSource is called when production uniqueJoinTableSource is entered.
func (s *BaseHiveParserListener) EnterUniqueJoinTableSource(ctx *UniqueJoinTableSourceContext) {}

// ExitUniqueJoinTableSource is called when production uniqueJoinTableSource is exited.
func (s *BaseHiveParserListener) ExitUniqueJoinTableSource(ctx *UniqueJoinTableSourceContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseHiveParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseHiveParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseHiveParserListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseHiveParserListener) ExitViewName(ctx *ViewNameContext) {}

// EnterSubQuerySource is called when production subQuerySource is entered.
func (s *BaseHiveParserListener) EnterSubQuerySource(ctx *SubQuerySourceContext) {}

// ExitSubQuerySource is called when production subQuerySource is exited.
func (s *BaseHiveParserListener) ExitSubQuerySource(ctx *SubQuerySourceContext) {}

// EnterPartitioningSpec is called when production partitioningSpec is entered.
func (s *BaseHiveParserListener) EnterPartitioningSpec(ctx *PartitioningSpecContext) {}

// ExitPartitioningSpec is called when production partitioningSpec is exited.
func (s *BaseHiveParserListener) ExitPartitioningSpec(ctx *PartitioningSpecContext) {}

// EnterPartitionTableFunctionSource is called when production partitionTableFunctionSource is entered.
func (s *BaseHiveParserListener) EnterPartitionTableFunctionSource(ctx *PartitionTableFunctionSourceContext) {
}

// ExitPartitionTableFunctionSource is called when production partitionTableFunctionSource is exited.
func (s *BaseHiveParserListener) ExitPartitionTableFunctionSource(ctx *PartitionTableFunctionSourceContext) {
}

// EnterPartitionedTableFunction is called when production partitionedTableFunction is entered.
func (s *BaseHiveParserListener) EnterPartitionedTableFunction(ctx *PartitionedTableFunctionContext) {
}

// ExitPartitionedTableFunction is called when production partitionedTableFunction is exited.
func (s *BaseHiveParserListener) ExitPartitionedTableFunction(ctx *PartitionedTableFunctionContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseHiveParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseHiveParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterSearchCondition is called when production searchCondition is entered.
func (s *BaseHiveParserListener) EnterSearchCondition(ctx *SearchConditionContext) {}

// ExitSearchCondition is called when production searchCondition is exited.
func (s *BaseHiveParserListener) ExitSearchCondition(ctx *SearchConditionContext) {}

// EnterValuesClause is called when production valuesClause is entered.
func (s *BaseHiveParserListener) EnterValuesClause(ctx *ValuesClauseContext) {}

// ExitValuesClause is called when production valuesClause is exited.
func (s *BaseHiveParserListener) ExitValuesClause(ctx *ValuesClauseContext) {}

// EnterValuesTableConstructor is called when production valuesTableConstructor is entered.
func (s *BaseHiveParserListener) EnterValuesTableConstructor(ctx *ValuesTableConstructorContext) {}

// ExitValuesTableConstructor is called when production valuesTableConstructor is exited.
func (s *BaseHiveParserListener) ExitValuesTableConstructor(ctx *ValuesTableConstructorContext) {}

// EnterValueRowConstructor is called when production valueRowConstructor is entered.
func (s *BaseHiveParserListener) EnterValueRowConstructor(ctx *ValueRowConstructorContext) {}

// ExitValueRowConstructor is called when production valueRowConstructor is exited.
func (s *BaseHiveParserListener) ExitValueRowConstructor(ctx *ValueRowConstructorContext) {}

// EnterVirtualTableSource is called when production virtualTableSource is entered.
func (s *BaseHiveParserListener) EnterVirtualTableSource(ctx *VirtualTableSourceContext) {}

// ExitVirtualTableSource is called when production virtualTableSource is exited.
func (s *BaseHiveParserListener) ExitVirtualTableSource(ctx *VirtualTableSourceContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseHiveParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseHiveParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupby_expression is called when production groupby_expression is entered.
func (s *BaseHiveParserListener) EnterGroupby_expression(ctx *Groupby_expressionContext) {}

// ExitGroupby_expression is called when production groupby_expression is exited.
func (s *BaseHiveParserListener) ExitGroupby_expression(ctx *Groupby_expressionContext) {}

// EnterGroupByEmpty is called when production groupByEmpty is entered.
func (s *BaseHiveParserListener) EnterGroupByEmpty(ctx *GroupByEmptyContext) {}

// ExitGroupByEmpty is called when production groupByEmpty is exited.
func (s *BaseHiveParserListener) ExitGroupByEmpty(ctx *GroupByEmptyContext) {}

// EnterRollupStandard is called when production rollupStandard is entered.
func (s *BaseHiveParserListener) EnterRollupStandard(ctx *RollupStandardContext) {}

// ExitRollupStandard is called when production rollupStandard is exited.
func (s *BaseHiveParserListener) ExitRollupStandard(ctx *RollupStandardContext) {}

// EnterRollupOldSyntax is called when production rollupOldSyntax is entered.
func (s *BaseHiveParserListener) EnterRollupOldSyntax(ctx *RollupOldSyntaxContext) {}

// ExitRollupOldSyntax is called when production rollupOldSyntax is exited.
func (s *BaseHiveParserListener) ExitRollupOldSyntax(ctx *RollupOldSyntaxContext) {}

// EnterGroupingSetExpression is called when production groupingSetExpression is entered.
func (s *BaseHiveParserListener) EnterGroupingSetExpression(ctx *GroupingSetExpressionContext) {}

// ExitGroupingSetExpression is called when production groupingSetExpression is exited.
func (s *BaseHiveParserListener) ExitGroupingSetExpression(ctx *GroupingSetExpressionContext) {}

// EnterGroupingSetExpressionMultiple is called when production groupingSetExpressionMultiple is entered.
func (s *BaseHiveParserListener) EnterGroupingSetExpressionMultiple(ctx *GroupingSetExpressionMultipleContext) {
}

// ExitGroupingSetExpressionMultiple is called when production groupingSetExpressionMultiple is exited.
func (s *BaseHiveParserListener) ExitGroupingSetExpressionMultiple(ctx *GroupingSetExpressionMultipleContext) {
}

// EnterGroupingExpressionSingle is called when production groupingExpressionSingle is entered.
func (s *BaseHiveParserListener) EnterGroupingExpressionSingle(ctx *GroupingExpressionSingleContext) {
}

// ExitGroupingExpressionSingle is called when production groupingExpressionSingle is exited.
func (s *BaseHiveParserListener) ExitGroupingExpressionSingle(ctx *GroupingExpressionSingleContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseHiveParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseHiveParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterHavingCondition is called when production havingCondition is entered.
func (s *BaseHiveParserListener) EnterHavingCondition(ctx *HavingConditionContext) {}

// ExitHavingCondition is called when production havingCondition is exited.
func (s *BaseHiveParserListener) ExitHavingCondition(ctx *HavingConditionContext) {}

// EnterExpressionsInParenthesis is called when production expressionsInParenthesis is entered.
func (s *BaseHiveParserListener) EnterExpressionsInParenthesis(ctx *ExpressionsInParenthesisContext) {
}

// ExitExpressionsInParenthesis is called when production expressionsInParenthesis is exited.
func (s *BaseHiveParserListener) ExitExpressionsInParenthesis(ctx *ExpressionsInParenthesisContext) {}

// EnterExpressionsNotInParenthesis is called when production expressionsNotInParenthesis is entered.
func (s *BaseHiveParserListener) EnterExpressionsNotInParenthesis(ctx *ExpressionsNotInParenthesisContext) {
}

// ExitExpressionsNotInParenthesis is called when production expressionsNotInParenthesis is exited.
func (s *BaseHiveParserListener) ExitExpressionsNotInParenthesis(ctx *ExpressionsNotInParenthesisContext) {
}

// EnterExpressionPart is called when production expressionPart is entered.
func (s *BaseHiveParserListener) EnterExpressionPart(ctx *ExpressionPartContext) {}

// ExitExpressionPart is called when production expressionPart is exited.
func (s *BaseHiveParserListener) ExitExpressionPart(ctx *ExpressionPartContext) {}

// EnterExpressions is called when production expressions is entered.
func (s *BaseHiveParserListener) EnterExpressions(ctx *ExpressionsContext) {}

// ExitExpressions is called when production expressions is exited.
func (s *BaseHiveParserListener) ExitExpressions(ctx *ExpressionsContext) {}

// EnterColumnRefOrderInParenthesis is called when production columnRefOrderInParenthesis is entered.
func (s *BaseHiveParserListener) EnterColumnRefOrderInParenthesis(ctx *ColumnRefOrderInParenthesisContext) {
}

// ExitColumnRefOrderInParenthesis is called when production columnRefOrderInParenthesis is exited.
func (s *BaseHiveParserListener) ExitColumnRefOrderInParenthesis(ctx *ColumnRefOrderInParenthesisContext) {
}

// EnterColumnRefOrderNotInParenthesis is called when production columnRefOrderNotInParenthesis is entered.
func (s *BaseHiveParserListener) EnterColumnRefOrderNotInParenthesis(ctx *ColumnRefOrderNotInParenthesisContext) {
}

// ExitColumnRefOrderNotInParenthesis is called when production columnRefOrderNotInParenthesis is exited.
func (s *BaseHiveParserListener) ExitColumnRefOrderNotInParenthesis(ctx *ColumnRefOrderNotInParenthesisContext) {
}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseHiveParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseHiveParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterClusterByClause is called when production clusterByClause is entered.
func (s *BaseHiveParserListener) EnterClusterByClause(ctx *ClusterByClauseContext) {}

// ExitClusterByClause is called when production clusterByClause is exited.
func (s *BaseHiveParserListener) ExitClusterByClause(ctx *ClusterByClauseContext) {}

// EnterPartitionByClause is called when production partitionByClause is entered.
func (s *BaseHiveParserListener) EnterPartitionByClause(ctx *PartitionByClauseContext) {}

// ExitPartitionByClause is called when production partitionByClause is exited.
func (s *BaseHiveParserListener) ExitPartitionByClause(ctx *PartitionByClauseContext) {}

// EnterDistributeByClause is called when production distributeByClause is entered.
func (s *BaseHiveParserListener) EnterDistributeByClause(ctx *DistributeByClauseContext) {}

// ExitDistributeByClause is called when production distributeByClause is exited.
func (s *BaseHiveParserListener) ExitDistributeByClause(ctx *DistributeByClauseContext) {}

// EnterSortByClause is called when production sortByClause is entered.
func (s *BaseHiveParserListener) EnterSortByClause(ctx *SortByClauseContext) {}

// ExitSortByClause is called when production sortByClause is exited.
func (s *BaseHiveParserListener) ExitSortByClause(ctx *SortByClauseContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseHiveParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseHiveParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseHiveParserListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseHiveParserListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterCastExpression is called when production castExpression is entered.
func (s *BaseHiveParserListener) EnterCastExpression(ctx *CastExpressionContext) {}

// ExitCastExpression is called when production castExpression is exited.
func (s *BaseHiveParserListener) ExitCastExpression(ctx *CastExpressionContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseHiveParserListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseHiveParserListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterWhenExpression is called when production whenExpression is entered.
func (s *BaseHiveParserListener) EnterWhenExpression(ctx *WhenExpressionContext) {}

// ExitWhenExpression is called when production whenExpression is exited.
func (s *BaseHiveParserListener) ExitWhenExpression(ctx *WhenExpressionContext) {}

// EnterFloorExpression is called when production floorExpression is entered.
func (s *BaseHiveParserListener) EnterFloorExpression(ctx *FloorExpressionContext) {}

// ExitFloorExpression is called when production floorExpression is exited.
func (s *BaseHiveParserListener) ExitFloorExpression(ctx *FloorExpressionContext) {}

// EnterFloorDateQualifiers is called when production floorDateQualifiers is entered.
func (s *BaseHiveParserListener) EnterFloorDateQualifiers(ctx *FloorDateQualifiersContext) {}

// ExitFloorDateQualifiers is called when production floorDateQualifiers is exited.
func (s *BaseHiveParserListener) ExitFloorDateQualifiers(ctx *FloorDateQualifiersContext) {}

// EnterExtractExpression is called when production extractExpression is entered.
func (s *BaseHiveParserListener) EnterExtractExpression(ctx *ExtractExpressionContext) {}

// ExitExtractExpression is called when production extractExpression is exited.
func (s *BaseHiveParserListener) ExitExtractExpression(ctx *ExtractExpressionContext) {}

// EnterTimeQualifiers is called when production timeQualifiers is entered.
func (s *BaseHiveParserListener) EnterTimeQualifiers(ctx *TimeQualifiersContext) {}

// ExitTimeQualifiers is called when production timeQualifiers is exited.
func (s *BaseHiveParserListener) ExitTimeQualifiers(ctx *TimeQualifiersContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseHiveParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseHiveParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterStringLiteralSequence is called when production stringLiteralSequence is entered.
func (s *BaseHiveParserListener) EnterStringLiteralSequence(ctx *StringLiteralSequenceContext) {}

// ExitStringLiteralSequence is called when production stringLiteralSequence is exited.
func (s *BaseHiveParserListener) ExitStringLiteralSequence(ctx *StringLiteralSequenceContext) {}

// EnterCharSetStringLiteral is called when production charSetStringLiteral is entered.
func (s *BaseHiveParserListener) EnterCharSetStringLiteral(ctx *CharSetStringLiteralContext) {}

// ExitCharSetStringLiteral is called when production charSetStringLiteral is exited.
func (s *BaseHiveParserListener) ExitCharSetStringLiteral(ctx *CharSetStringLiteralContext) {}

// EnterDateLiteral is called when production dateLiteral is entered.
func (s *BaseHiveParserListener) EnterDateLiteral(ctx *DateLiteralContext) {}

// ExitDateLiteral is called when production dateLiteral is exited.
func (s *BaseHiveParserListener) ExitDateLiteral(ctx *DateLiteralContext) {}

// EnterTimestampLiteral is called when production timestampLiteral is entered.
func (s *BaseHiveParserListener) EnterTimestampLiteral(ctx *TimestampLiteralContext) {}

// ExitTimestampLiteral is called when production timestampLiteral is exited.
func (s *BaseHiveParserListener) ExitTimestampLiteral(ctx *TimestampLiteralContext) {}

// EnterTimestampLocalTZLiteral is called when production timestampLocalTZLiteral is entered.
func (s *BaseHiveParserListener) EnterTimestampLocalTZLiteral(ctx *TimestampLocalTZLiteralContext) {}

// ExitTimestampLocalTZLiteral is called when production timestampLocalTZLiteral is exited.
func (s *BaseHiveParserListener) ExitTimestampLocalTZLiteral(ctx *TimestampLocalTZLiteralContext) {}

// EnterIntervalValue is called when production intervalValue is entered.
func (s *BaseHiveParserListener) EnterIntervalValue(ctx *IntervalValueContext) {}

// ExitIntervalValue is called when production intervalValue is exited.
func (s *BaseHiveParserListener) ExitIntervalValue(ctx *IntervalValueContext) {}

// EnterIntervalLiteral is called when production intervalLiteral is entered.
func (s *BaseHiveParserListener) EnterIntervalLiteral(ctx *IntervalLiteralContext) {}

// ExitIntervalLiteral is called when production intervalLiteral is exited.
func (s *BaseHiveParserListener) ExitIntervalLiteral(ctx *IntervalLiteralContext) {}

// EnterIntervalExpression is called when production intervalExpression is entered.
func (s *BaseHiveParserListener) EnterIntervalExpression(ctx *IntervalExpressionContext) {}

// ExitIntervalExpression is called when production intervalExpression is exited.
func (s *BaseHiveParserListener) ExitIntervalExpression(ctx *IntervalExpressionContext) {}

// EnterIntervalQualifiers is called when production intervalQualifiers is entered.
func (s *BaseHiveParserListener) EnterIntervalQualifiers(ctx *IntervalQualifiersContext) {}

// ExitIntervalQualifiers is called when production intervalQualifiers is exited.
func (s *BaseHiveParserListener) ExitIntervalQualifiers(ctx *IntervalQualifiersContext) {}

// EnterAtomExpression is called when production atomExpression is entered.
func (s *BaseHiveParserListener) EnterAtomExpression(ctx *AtomExpressionContext) {}

// ExitAtomExpression is called when production atomExpression is exited.
func (s *BaseHiveParserListener) ExitAtomExpression(ctx *AtomExpressionContext) {}

// EnterPrecedenceUnaryOperator is called when production precedenceUnaryOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceUnaryOperator(ctx *PrecedenceUnaryOperatorContext) {}

// ExitPrecedenceUnaryOperator is called when production precedenceUnaryOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceUnaryOperator(ctx *PrecedenceUnaryOperatorContext) {}

// EnterIsCondition is called when production isCondition is entered.
func (s *BaseHiveParserListener) EnterIsCondition(ctx *IsConditionContext) {}

// ExitIsCondition is called when production isCondition is exited.
func (s *BaseHiveParserListener) ExitIsCondition(ctx *IsConditionContext) {}

// EnterPrecedenceBitwiseXorOperator is called when production precedenceBitwiseXorOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceBitwiseXorOperator(ctx *PrecedenceBitwiseXorOperatorContext) {
}

// ExitPrecedenceBitwiseXorOperator is called when production precedenceBitwiseXorOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceBitwiseXorOperator(ctx *PrecedenceBitwiseXorOperatorContext) {
}

// EnterPrecedenceStarOperator is called when production precedenceStarOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceStarOperator(ctx *PrecedenceStarOperatorContext) {}

// ExitPrecedenceStarOperator is called when production precedenceStarOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceStarOperator(ctx *PrecedenceStarOperatorContext) {}

// EnterPrecedencePlusOperator is called when production precedencePlusOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedencePlusOperator(ctx *PrecedencePlusOperatorContext) {}

// ExitPrecedencePlusOperator is called when production precedencePlusOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedencePlusOperator(ctx *PrecedencePlusOperatorContext) {}

// EnterPrecedenceConcatenateOperator is called when production precedenceConcatenateOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceConcatenateOperator(ctx *PrecedenceConcatenateOperatorContext) {
}

// ExitPrecedenceConcatenateOperator is called when production precedenceConcatenateOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceConcatenateOperator(ctx *PrecedenceConcatenateOperatorContext) {
}

// EnterPrecedenceAmpersandOperator is called when production precedenceAmpersandOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceAmpersandOperator(ctx *PrecedenceAmpersandOperatorContext) {
}

// ExitPrecedenceAmpersandOperator is called when production precedenceAmpersandOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceAmpersandOperator(ctx *PrecedenceAmpersandOperatorContext) {
}

// EnterPrecedenceBitwiseOrOperator is called when production precedenceBitwiseOrOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceBitwiseOrOperator(ctx *PrecedenceBitwiseOrOperatorContext) {
}

// ExitPrecedenceBitwiseOrOperator is called when production precedenceBitwiseOrOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceBitwiseOrOperator(ctx *PrecedenceBitwiseOrOperatorContext) {
}

// EnterPrecedenceRegexpOperator is called when production precedenceRegexpOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceRegexpOperator(ctx *PrecedenceRegexpOperatorContext) {
}

// ExitPrecedenceRegexpOperator is called when production precedenceRegexpOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceRegexpOperator(ctx *PrecedenceRegexpOperatorContext) {}

// EnterPrecedenceSimilarOperator is called when production precedenceSimilarOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceSimilarOperator(ctx *PrecedenceSimilarOperatorContext) {
}

// ExitPrecedenceSimilarOperator is called when production precedenceSimilarOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceSimilarOperator(ctx *PrecedenceSimilarOperatorContext) {
}

// EnterPrecedenceDistinctOperator is called when production precedenceDistinctOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceDistinctOperator(ctx *PrecedenceDistinctOperatorContext) {
}

// ExitPrecedenceDistinctOperator is called when production precedenceDistinctOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceDistinctOperator(ctx *PrecedenceDistinctOperatorContext) {
}

// EnterPrecedenceEqualOperator is called when production precedenceEqualOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceEqualOperator(ctx *PrecedenceEqualOperatorContext) {}

// ExitPrecedenceEqualOperator is called when production precedenceEqualOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceEqualOperator(ctx *PrecedenceEqualOperatorContext) {}

// EnterPrecedenceNotOperator is called when production precedenceNotOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceNotOperator(ctx *PrecedenceNotOperatorContext) {}

// ExitPrecedenceNotOperator is called when production precedenceNotOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceNotOperator(ctx *PrecedenceNotOperatorContext) {}

// EnterPrecedenceAndOperator is called when production precedenceAndOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceAndOperator(ctx *PrecedenceAndOperatorContext) {}

// ExitPrecedenceAndOperator is called when production precedenceAndOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceAndOperator(ctx *PrecedenceAndOperatorContext) {}

// EnterPrecedenceOrOperator is called when production precedenceOrOperator is entered.
func (s *BaseHiveParserListener) EnterPrecedenceOrOperator(ctx *PrecedenceOrOperatorContext) {}

// ExitPrecedenceOrOperator is called when production precedenceOrOperator is exited.
func (s *BaseHiveParserListener) ExitPrecedenceOrOperator(ctx *PrecedenceOrOperatorContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHiveParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHiveParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSubQueryExpression is called when production subQueryExpression is entered.
func (s *BaseHiveParserListener) EnterSubQueryExpression(ctx *SubQueryExpressionContext) {}

// ExitSubQueryExpression is called when production subQueryExpression is exited.
func (s *BaseHiveParserListener) ExitSubQueryExpression(ctx *SubQueryExpressionContext) {}

// EnterPrecedenceSimilarExpressionPart is called when production precedenceSimilarExpressionPart is entered.
func (s *BaseHiveParserListener) EnterPrecedenceSimilarExpressionPart(ctx *PrecedenceSimilarExpressionPartContext) {
}

// ExitPrecedenceSimilarExpressionPart is called when production precedenceSimilarExpressionPart is exited.
func (s *BaseHiveParserListener) ExitPrecedenceSimilarExpressionPart(ctx *PrecedenceSimilarExpressionPartContext) {
}

// EnterPrecedenceSimilarExpressionAtom is called when production precedenceSimilarExpressionAtom is entered.
func (s *BaseHiveParserListener) EnterPrecedenceSimilarExpressionAtom(ctx *PrecedenceSimilarExpressionAtomContext) {
}

// ExitPrecedenceSimilarExpressionAtom is called when production precedenceSimilarExpressionAtom is exited.
func (s *BaseHiveParserListener) ExitPrecedenceSimilarExpressionAtom(ctx *PrecedenceSimilarExpressionAtomContext) {
}

// EnterPrecedenceSimilarExpressionIn is called when production precedenceSimilarExpressionIn is entered.
func (s *BaseHiveParserListener) EnterPrecedenceSimilarExpressionIn(ctx *PrecedenceSimilarExpressionInContext) {
}

// ExitPrecedenceSimilarExpressionIn is called when production precedenceSimilarExpressionIn is exited.
func (s *BaseHiveParserListener) ExitPrecedenceSimilarExpressionIn(ctx *PrecedenceSimilarExpressionInContext) {
}

// EnterPrecedenceSimilarExpressionPartNot is called when production precedenceSimilarExpressionPartNot is entered.
func (s *BaseHiveParserListener) EnterPrecedenceSimilarExpressionPartNot(ctx *PrecedenceSimilarExpressionPartNotContext) {
}

// ExitPrecedenceSimilarExpressionPartNot is called when production precedenceSimilarExpressionPartNot is exited.
func (s *BaseHiveParserListener) ExitPrecedenceSimilarExpressionPartNot(ctx *PrecedenceSimilarExpressionPartNotContext) {
}

// EnterBooleanValue is called when production booleanValue is entered.
func (s *BaseHiveParserListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production booleanValue is exited.
func (s *BaseHiveParserListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterBooleanValueTok is called when production booleanValueTok is entered.
func (s *BaseHiveParserListener) EnterBooleanValueTok(ctx *BooleanValueTokContext) {}

// ExitBooleanValueTok is called when production booleanValueTok is exited.
func (s *BaseHiveParserListener) ExitBooleanValueTok(ctx *BooleanValueTokContext) {}

// EnterTableOrPartition is called when production tableOrPartition is entered.
func (s *BaseHiveParserListener) EnterTableOrPartition(ctx *TableOrPartitionContext) {}

// ExitTableOrPartition is called when production tableOrPartition is exited.
func (s *BaseHiveParserListener) ExitTableOrPartition(ctx *TableOrPartitionContext) {}

// EnterPartitionSpec is called when production partitionSpec is entered.
func (s *BaseHiveParserListener) EnterPartitionSpec(ctx *PartitionSpecContext) {}

// ExitPartitionSpec is called when production partitionSpec is exited.
func (s *BaseHiveParserListener) ExitPartitionSpec(ctx *PartitionSpecContext) {}

// EnterPartitionVal is called when production partitionVal is entered.
func (s *BaseHiveParserListener) EnterPartitionVal(ctx *PartitionValContext) {}

// ExitPartitionVal is called when production partitionVal is exited.
func (s *BaseHiveParserListener) ExitPartitionVal(ctx *PartitionValContext) {}

// EnterDropPartitionSpec is called when production dropPartitionSpec is entered.
func (s *BaseHiveParserListener) EnterDropPartitionSpec(ctx *DropPartitionSpecContext) {}

// ExitDropPartitionSpec is called when production dropPartitionSpec is exited.
func (s *BaseHiveParserListener) ExitDropPartitionSpec(ctx *DropPartitionSpecContext) {}

// EnterDropPartitionVal is called when production dropPartitionVal is entered.
func (s *BaseHiveParserListener) EnterDropPartitionVal(ctx *DropPartitionValContext) {}

// ExitDropPartitionVal is called when production dropPartitionVal is exited.
func (s *BaseHiveParserListener) ExitDropPartitionVal(ctx *DropPartitionValContext) {}

// EnterDropPartitionOperator is called when production dropPartitionOperator is entered.
func (s *BaseHiveParserListener) EnterDropPartitionOperator(ctx *DropPartitionOperatorContext) {}

// ExitDropPartitionOperator is called when production dropPartitionOperator is exited.
func (s *BaseHiveParserListener) ExitDropPartitionOperator(ctx *DropPartitionOperatorContext) {}

// EnterSysFuncNames is called when production sysFuncNames is entered.
func (s *BaseHiveParserListener) EnterSysFuncNames(ctx *SysFuncNamesContext) {}

// ExitSysFuncNames is called when production sysFuncNames is exited.
func (s *BaseHiveParserListener) ExitSysFuncNames(ctx *SysFuncNamesContext) {}

// EnterDescFuncNames is called when production descFuncNames is entered.
func (s *BaseHiveParserListener) EnterDescFuncNames(ctx *DescFuncNamesContext) {}

// ExitDescFuncNames is called when production descFuncNames is exited.
func (s *BaseHiveParserListener) ExitDescFuncNames(ctx *DescFuncNamesContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseHiveParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseHiveParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterFunctionIdentifier is called when production functionIdentifier is entered.
func (s *BaseHiveParserListener) EnterFunctionIdentifier(ctx *FunctionIdentifierContext) {}

// ExitFunctionIdentifier is called when production functionIdentifier is exited.
func (s *BaseHiveParserListener) ExitFunctionIdentifier(ctx *FunctionIdentifierContext) {}

// EnterPrincipalIdentifier is called when production principalIdentifier is entered.
func (s *BaseHiveParserListener) EnterPrincipalIdentifier(ctx *PrincipalIdentifierContext) {}

// ExitPrincipalIdentifier is called when production principalIdentifier is exited.
func (s *BaseHiveParserListener) ExitPrincipalIdentifier(ctx *PrincipalIdentifierContext) {}

// EnterNonReserved is called when production nonReserved is entered.
func (s *BaseHiveParserListener) EnterNonReserved(ctx *NonReservedContext) {}

// ExitNonReserved is called when production nonReserved is exited.
func (s *BaseHiveParserListener) ExitNonReserved(ctx *NonReservedContext) {}

// EnterSql11ReservedKeywordsUsedAsFunctionName is called when production sql11ReservedKeywordsUsedAsFunctionName is entered.
func (s *BaseHiveParserListener) EnterSql11ReservedKeywordsUsedAsFunctionName(ctx *Sql11ReservedKeywordsUsedAsFunctionNameContext) {
}

// ExitSql11ReservedKeywordsUsedAsFunctionName is called when production sql11ReservedKeywordsUsedAsFunctionName is exited.
func (s *BaseHiveParserListener) ExitSql11ReservedKeywordsUsedAsFunctionName(ctx *Sql11ReservedKeywordsUsedAsFunctionNameContext) {
}

// EnterResourcePlanDdlStatements is called when production resourcePlanDdlStatements is entered.
func (s *BaseHiveParserListener) EnterResourcePlanDdlStatements(ctx *ResourcePlanDdlStatementsContext) {
}

// ExitResourcePlanDdlStatements is called when production resourcePlanDdlStatements is exited.
func (s *BaseHiveParserListener) ExitResourcePlanDdlStatements(ctx *ResourcePlanDdlStatementsContext) {
}

// EnterRpAssign is called when production rpAssign is entered.
func (s *BaseHiveParserListener) EnterRpAssign(ctx *RpAssignContext) {}

// ExitRpAssign is called when production rpAssign is exited.
func (s *BaseHiveParserListener) ExitRpAssign(ctx *RpAssignContext) {}

// EnterRpAssignList is called when production rpAssignList is entered.
func (s *BaseHiveParserListener) EnterRpAssignList(ctx *RpAssignListContext) {}

// ExitRpAssignList is called when production rpAssignList is exited.
func (s *BaseHiveParserListener) ExitRpAssignList(ctx *RpAssignListContext) {}

// EnterRpUnassign is called when production rpUnassign is entered.
func (s *BaseHiveParserListener) EnterRpUnassign(ctx *RpUnassignContext) {}

// ExitRpUnassign is called when production rpUnassign is exited.
func (s *BaseHiveParserListener) ExitRpUnassign(ctx *RpUnassignContext) {}

// EnterRpUnassignList is called when production rpUnassignList is entered.
func (s *BaseHiveParserListener) EnterRpUnassignList(ctx *RpUnassignListContext) {}

// ExitRpUnassignList is called when production rpUnassignList is exited.
func (s *BaseHiveParserListener) ExitRpUnassignList(ctx *RpUnassignListContext) {}

// EnterCreateResourcePlanStatement is called when production createResourcePlanStatement is entered.
func (s *BaseHiveParserListener) EnterCreateResourcePlanStatement(ctx *CreateResourcePlanStatementContext) {
}

// ExitCreateResourcePlanStatement is called when production createResourcePlanStatement is exited.
func (s *BaseHiveParserListener) ExitCreateResourcePlanStatement(ctx *CreateResourcePlanStatementContext) {
}

// EnterWithReplace is called when production withReplace is entered.
func (s *BaseHiveParserListener) EnterWithReplace(ctx *WithReplaceContext) {}

// ExitWithReplace is called when production withReplace is exited.
func (s *BaseHiveParserListener) ExitWithReplace(ctx *WithReplaceContext) {}

// EnterActivate is called when production activate is entered.
func (s *BaseHiveParserListener) EnterActivate(ctx *ActivateContext) {}

// ExitActivate is called when production activate is exited.
func (s *BaseHiveParserListener) ExitActivate(ctx *ActivateContext) {}

// EnterEnable is called when production enable is entered.
func (s *BaseHiveParserListener) EnterEnable(ctx *EnableContext) {}

// ExitEnable is called when production enable is exited.
func (s *BaseHiveParserListener) ExitEnable(ctx *EnableContext) {}

// EnterDisable is called when production disable is entered.
func (s *BaseHiveParserListener) EnterDisable(ctx *DisableContext) {}

// ExitDisable is called when production disable is exited.
func (s *BaseHiveParserListener) ExitDisable(ctx *DisableContext) {}

// EnterUnmanaged is called when production unmanaged is entered.
func (s *BaseHiveParserListener) EnterUnmanaged(ctx *UnmanagedContext) {}

// ExitUnmanaged is called when production unmanaged is exited.
func (s *BaseHiveParserListener) ExitUnmanaged(ctx *UnmanagedContext) {}

// EnterAlterResourcePlanStatement is called when production alterResourcePlanStatement is entered.
func (s *BaseHiveParserListener) EnterAlterResourcePlanStatement(ctx *AlterResourcePlanStatementContext) {
}

// ExitAlterResourcePlanStatement is called when production alterResourcePlanStatement is exited.
func (s *BaseHiveParserListener) ExitAlterResourcePlanStatement(ctx *AlterResourcePlanStatementContext) {
}

// EnterGlobalWmStatement is called when production globalWmStatement is entered.
func (s *BaseHiveParserListener) EnterGlobalWmStatement(ctx *GlobalWmStatementContext) {}

// ExitGlobalWmStatement is called when production globalWmStatement is exited.
func (s *BaseHiveParserListener) ExitGlobalWmStatement(ctx *GlobalWmStatementContext) {}

// EnterReplaceResourcePlanStatement is called when production replaceResourcePlanStatement is entered.
func (s *BaseHiveParserListener) EnterReplaceResourcePlanStatement(ctx *ReplaceResourcePlanStatementContext) {
}

// ExitReplaceResourcePlanStatement is called when production replaceResourcePlanStatement is exited.
func (s *BaseHiveParserListener) ExitReplaceResourcePlanStatement(ctx *ReplaceResourcePlanStatementContext) {
}

// EnterDropResourcePlanStatement is called when production dropResourcePlanStatement is entered.
func (s *BaseHiveParserListener) EnterDropResourcePlanStatement(ctx *DropResourcePlanStatementContext) {
}

// ExitDropResourcePlanStatement is called when production dropResourcePlanStatement is exited.
func (s *BaseHiveParserListener) ExitDropResourcePlanStatement(ctx *DropResourcePlanStatementContext) {
}

// EnterPoolPath is called when production poolPath is entered.
func (s *BaseHiveParserListener) EnterPoolPath(ctx *PoolPathContext) {}

// ExitPoolPath is called when production poolPath is exited.
func (s *BaseHiveParserListener) ExitPoolPath(ctx *PoolPathContext) {}

// EnterTriggerExpression is called when production triggerExpression is entered.
func (s *BaseHiveParserListener) EnterTriggerExpression(ctx *TriggerExpressionContext) {}

// ExitTriggerExpression is called when production triggerExpression is exited.
func (s *BaseHiveParserListener) ExitTriggerExpression(ctx *TriggerExpressionContext) {}

// EnterTriggerExpressionStandalone is called when production triggerExpressionStandalone is entered.
func (s *BaseHiveParserListener) EnterTriggerExpressionStandalone(ctx *TriggerExpressionStandaloneContext) {
}

// ExitTriggerExpressionStandalone is called when production triggerExpressionStandalone is exited.
func (s *BaseHiveParserListener) ExitTriggerExpressionStandalone(ctx *TriggerExpressionStandaloneContext) {
}

// EnterTriggerOrExpression is called when production triggerOrExpression is entered.
func (s *BaseHiveParserListener) EnterTriggerOrExpression(ctx *TriggerOrExpressionContext) {}

// ExitTriggerOrExpression is called when production triggerOrExpression is exited.
func (s *BaseHiveParserListener) ExitTriggerOrExpression(ctx *TriggerOrExpressionContext) {}

// EnterTriggerAndExpression is called when production triggerAndExpression is entered.
func (s *BaseHiveParserListener) EnterTriggerAndExpression(ctx *TriggerAndExpressionContext) {}

// ExitTriggerAndExpression is called when production triggerAndExpression is exited.
func (s *BaseHiveParserListener) ExitTriggerAndExpression(ctx *TriggerAndExpressionContext) {}

// EnterTriggerAtomExpression is called when production triggerAtomExpression is entered.
func (s *BaseHiveParserListener) EnterTriggerAtomExpression(ctx *TriggerAtomExpressionContext) {}

// ExitTriggerAtomExpression is called when production triggerAtomExpression is exited.
func (s *BaseHiveParserListener) ExitTriggerAtomExpression(ctx *TriggerAtomExpressionContext) {}

// EnterTriggerLiteral is called when production triggerLiteral is entered.
func (s *BaseHiveParserListener) EnterTriggerLiteral(ctx *TriggerLiteralContext) {}

// ExitTriggerLiteral is called when production triggerLiteral is exited.
func (s *BaseHiveParserListener) ExitTriggerLiteral(ctx *TriggerLiteralContext) {}

// EnterComparisionOperator is called when production comparisionOperator is entered.
func (s *BaseHiveParserListener) EnterComparisionOperator(ctx *ComparisionOperatorContext) {}

// ExitComparisionOperator is called when production comparisionOperator is exited.
func (s *BaseHiveParserListener) ExitComparisionOperator(ctx *ComparisionOperatorContext) {}

// EnterTriggerActionExpression is called when production triggerActionExpression is entered.
func (s *BaseHiveParserListener) EnterTriggerActionExpression(ctx *TriggerActionExpressionContext) {}

// ExitTriggerActionExpression is called when production triggerActionExpression is exited.
func (s *BaseHiveParserListener) ExitTriggerActionExpression(ctx *TriggerActionExpressionContext) {}

// EnterTriggerActionExpressionStandalone is called when production triggerActionExpressionStandalone is entered.
func (s *BaseHiveParserListener) EnterTriggerActionExpressionStandalone(ctx *TriggerActionExpressionStandaloneContext) {
}

// ExitTriggerActionExpressionStandalone is called when production triggerActionExpressionStandalone is exited.
func (s *BaseHiveParserListener) ExitTriggerActionExpressionStandalone(ctx *TriggerActionExpressionStandaloneContext) {
}

// EnterCreateTriggerStatement is called when production createTriggerStatement is entered.
func (s *BaseHiveParserListener) EnterCreateTriggerStatement(ctx *CreateTriggerStatementContext) {}

// ExitCreateTriggerStatement is called when production createTriggerStatement is exited.
func (s *BaseHiveParserListener) ExitCreateTriggerStatement(ctx *CreateTriggerStatementContext) {}

// EnterAlterTriggerStatement is called when production alterTriggerStatement is entered.
func (s *BaseHiveParserListener) EnterAlterTriggerStatement(ctx *AlterTriggerStatementContext) {}

// ExitAlterTriggerStatement is called when production alterTriggerStatement is exited.
func (s *BaseHiveParserListener) ExitAlterTriggerStatement(ctx *AlterTriggerStatementContext) {}

// EnterDropTriggerStatement is called when production dropTriggerStatement is entered.
func (s *BaseHiveParserListener) EnterDropTriggerStatement(ctx *DropTriggerStatementContext) {}

// ExitDropTriggerStatement is called when production dropTriggerStatement is exited.
func (s *BaseHiveParserListener) ExitDropTriggerStatement(ctx *DropTriggerStatementContext) {}

// EnterPoolAssign is called when production poolAssign is entered.
func (s *BaseHiveParserListener) EnterPoolAssign(ctx *PoolAssignContext) {}

// ExitPoolAssign is called when production poolAssign is exited.
func (s *BaseHiveParserListener) ExitPoolAssign(ctx *PoolAssignContext) {}

// EnterPoolAssignList is called when production poolAssignList is entered.
func (s *BaseHiveParserListener) EnterPoolAssignList(ctx *PoolAssignListContext) {}

// ExitPoolAssignList is called when production poolAssignList is exited.
func (s *BaseHiveParserListener) ExitPoolAssignList(ctx *PoolAssignListContext) {}

// EnterCreatePoolStatement is called when production createPoolStatement is entered.
func (s *BaseHiveParserListener) EnterCreatePoolStatement(ctx *CreatePoolStatementContext) {}

// ExitCreatePoolStatement is called when production createPoolStatement is exited.
func (s *BaseHiveParserListener) ExitCreatePoolStatement(ctx *CreatePoolStatementContext) {}

// EnterAlterPoolStatement is called when production alterPoolStatement is entered.
func (s *BaseHiveParserListener) EnterAlterPoolStatement(ctx *AlterPoolStatementContext) {}

// ExitAlterPoolStatement is called when production alterPoolStatement is exited.
func (s *BaseHiveParserListener) ExitAlterPoolStatement(ctx *AlterPoolStatementContext) {}

// EnterDropPoolStatement is called when production dropPoolStatement is entered.
func (s *BaseHiveParserListener) EnterDropPoolStatement(ctx *DropPoolStatementContext) {}

// ExitDropPoolStatement is called when production dropPoolStatement is exited.
func (s *BaseHiveParserListener) ExitDropPoolStatement(ctx *DropPoolStatementContext) {}

// EnterCreateMappingStatement is called when production createMappingStatement is entered.
func (s *BaseHiveParserListener) EnterCreateMappingStatement(ctx *CreateMappingStatementContext) {}

// ExitCreateMappingStatement is called when production createMappingStatement is exited.
func (s *BaseHiveParserListener) ExitCreateMappingStatement(ctx *CreateMappingStatementContext) {}

// EnterAlterMappingStatement is called when production alterMappingStatement is entered.
func (s *BaseHiveParserListener) EnterAlterMappingStatement(ctx *AlterMappingStatementContext) {}

// ExitAlterMappingStatement is called when production alterMappingStatement is exited.
func (s *BaseHiveParserListener) ExitAlterMappingStatement(ctx *AlterMappingStatementContext) {}

// EnterDropMappingStatement is called when production dropMappingStatement is entered.
func (s *BaseHiveParserListener) EnterDropMappingStatement(ctx *DropMappingStatementContext) {}

// ExitDropMappingStatement is called when production dropMappingStatement is exited.
func (s *BaseHiveParserListener) ExitDropMappingStatement(ctx *DropMappingStatementContext) {}
