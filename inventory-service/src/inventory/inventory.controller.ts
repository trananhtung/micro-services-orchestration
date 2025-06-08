import { Controller, Get, Param, Query, Post, Body, BadRequestException } from '@nestjs/common';
import { InventoryService } from './inventory.service';
import { ApiTags, ApiOperation, ApiResponse, ApiBody, ApiParam } from '@nestjs/swagger';
import { Inventory } from './inventory.entity';

@ApiTags('Inventory')
@Controller('inventory')
export class InventoryController {
  constructor(private readonly inventoryService: InventoryService) {}

  @ApiOperation({ summary: 'Get all inventory items' })
  @ApiResponse({ status: 200, type: [Inventory] })
  @Get()
  findAll() {
    return this.inventoryService.findAll();
  }

  @ApiOperation({ summary: 'Get inventory item by ID' })
  @ApiParam({ name: 'id', type: Number })
  @ApiResponse({ status: 200, type: Inventory })
  @Get(':id')
  findOne(@Param('id') id: string) {
    return this.inventoryService.findOne(Number(id));
  }

  @ApiOperation({ summary: 'Check stock by product name' })
  @ApiParam({ name: 'productName', type: String })
  @ApiResponse({ status: 200, description: 'Quantity in stock', schema: { type: 'number', example: 100 } })
  @Get('/check-stock/:productName')
  checkStock(@Param('productName') productName: string) {
    return this.inventoryService.checkStock(productName);
  }

  @ApiOperation({ summary: 'Update or create inventory stock' })
  @ApiBody({
    schema: {
      type: 'object',
      properties: {
        productName: { type: 'string', example: 'Product A' },
        quantity: { type: 'number', example: 100 },
      },
      required: ['productName', 'quantity'],
    },
  })
  @ApiResponse({ status: 201, type: Inventory })
  @Post('/update-stock')
  updateStock(
    @Body('productName') productName: string,
    @Body('quantity') quantity: number,
  ) {
    if (!productName || quantity === undefined) {
      throw new BadRequestException('productName and quantity are required');
    }
    return this.inventoryService.updateStock(productName, quantity);
  }
}
